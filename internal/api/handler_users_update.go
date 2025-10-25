package api

import (
	"net/http"
	"encoding/json"

	"github.com/joseph-m-valdez/chirpy/internal/auth"
	"github.com/joseph-m-valdez/chirpy/internal/database"
)

func (a *API) HandlerUpdateUsersAuth(w http.ResponseWriter, req *http.Request) {
	type requestParams struct {
		Email	string			`json:"email"`
		Password	string	`json:"password"`
	}

   bearer, err := auth.GetBearerToken(req.Header)
	 if err != nil {
      respondWithError(w, http.StatusUnauthorized, "unauthorized", err)
		 	return
	 }

	 userID, err := auth.ValidateJWT(bearer, a.JWTSecret)

		if err != nil { 
        respondWithError(w, http.StatusUnauthorized, "unauthorized", err)
        return
		}

		params := requestParams{}
    if err = json.NewDecoder(req.Body).Decode(&params); err != nil {
        respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
        return
    }

		hashedPassword, err := auth.HashPassword(params.Password)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "error with pw", err)
			return
		}


		updatedUser, err := a.DB.UpdateUserAuth(req.Context(), database.UpdateUserAuthParams{
			Email: 					params.Email,
			HashedPassword: hashedPassword,
			ID:							userID,
		}) 

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "could not update user", err)
			return
		}

		respondWithJSON(w, http.StatusOK, User{
			ID: updatedUser.ID,	
			CreatedAt: updatedUser.CreatedAt,
			UpdatedAt:	updatedUser.UpdatedAt,
			Email:			updatedUser.Email,
		})
}
