package api

import (
	"encoding/json"
	"net/http"

	"github.com/joseph-m-valdez/chirpy/internal/database"
	"github.com/joseph-m-valdez/chirpy/internal/auth"
)

func (a *API) HandlerCreateUsers(w http.ResponseWriter, req *http.Request) {
	type createUserRequest struct {
		Password 	string	`json:"password"`
		Email			string	`json:"email"`
	}

	type createUserResponse User

	decoder := json.NewDecoder(req.Body)
	createUserReq := createUserRequest{}
	if err := decoder.Decode(&createUserReq); err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode createUserRequest", err)	
		return
	}	

	pw, err := auth.HashPassword(createUserReq.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create user", err)
		return
	}

	createdUser, err := a.DB.CreateUser(req.Context(), database.CreateUserParams{
		Email:					createUserReq.Email,
		HashedPassword:	pw,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create user", err)
		return
	}

	user := createUserResponse{
		ID: 				createdUser.ID,
		CreatedAt: 	createdUser.CreatedAt,
		UpdatedAt: 	createdUser.UpdatedAt,
		Email:		 	createdUser.Email,
	}

	respondWithJSON(w, http.StatusCreated, user )

}
