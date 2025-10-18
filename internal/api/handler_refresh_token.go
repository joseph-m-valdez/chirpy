package api

import (
	"net/http"
	"time"

	"github.com/joseph-m-valdez/chirpy/internal/auth"
)

func (a *API) HandlerRefreshToken(w http.ResponseWriter, req *http.Request) {
	 type response struct {
		 Token	string	`json:"token"`	
	 }

   bearer, err := auth.GetBearerToken(req.Header)
	 if err != nil {
			respondWithError(w, http.StatusInternalServerError, "no bearer token", err)
		 	return
	 }

	 user, err := a.DB.GetUserFromRefreshToken(req.Context(), bearer)
	 if err != nil {
			respondWithError(w, http.StatusUnauthorized, "unauthorized", err)
		 	return
	 }

	token, err := auth.MakeJWT(user.ID, a.JWTSecret, time.Duration(defaultJWTExpirySeconds)*time.Second)
		if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create token", err)
		return
	}


	 respondWithJSON(w, http.StatusOK, response{
		Token: token,
	 })

}
