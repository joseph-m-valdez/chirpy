package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerCreateUsers(w http.ResponseWriter, req *http.Request) {
	type createUserRequest struct {
		Email	string	`json:"email"`
	}

	type createUserResponse User

	decoder := json.NewDecoder(req.Body)
	createUserReq := createUserRequest{}
	if err := decoder.Decode(&createUserReq); err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode createUserRequest", err)	
		return
	}	

	createdUser, err := cfg.db.CreateUser(req.Context(), createUserReq.Email)
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
