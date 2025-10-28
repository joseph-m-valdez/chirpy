package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/joseph-m-valdez/chirpy/internal/auth"
	"github.com/joseph-m-valdez/chirpy/internal/database"
)

const defaultJWTExpirySeconds = 3600 // 1 hour

func (a *API) HandlerLogin(w http.ResponseWriter, req *http.Request) {
	type loginRequest struct {
		Password					string `json:"password"`
		Email							string	`json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	loginReq := loginRequest{}
	if err := decoder.Decode(&loginReq); err != nil {
		respondWithError(w, http.StatusInternalServerError, "error decoding params", err)
		return
	}

	user, err := a.DB.GetUser(req.Context(), loginReq.Email)
	if err != nil {	
		respondWithError(w, http.StatusInternalServerError, "error getting user", err)
		return
	}

	match, err := auth.CheckPasswordHash(loginReq.Password, user.HashedPassword)	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to login", err)
		return
	}

	if !match {
		respondWithError(w, http.StatusUnauthorized, "email or password is incorrect", err)
		return
	}

	token, err := auth.MakeJWT(user.ID, a.JWTSecret, time.Duration(defaultJWTExpirySeconds)*time.Second)
		if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create token", err)
		return
	}
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create refresh token", err)
		return
	}

	a.DB.CreateRefreshToken(req.Context(), database.CreateRefreshTokenParams{
			Token: 			refreshToken,
			UserID:			user.ID,
			ExpiresAt:	time.Now().UTC().Add(60 * 24 * time.Hour),
		},
	)


	respondWithJSON(w, http.StatusOK, User{
		ID:						user.ID,
		CreatedAt:		user.CreatedAt,
		UpdatedAt: 		user.UpdatedAt,
		Email:				user.Email,
		Token:				token,
		RefreshToken:	refreshToken,
		IsChirpyRed:	user.IsChirpyRed,
	})
}
