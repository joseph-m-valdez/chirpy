package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/joseph-m-valdez/chirpy/internal/database"
)

func (a *API) HandlerCreateChirps(w http.ResponseWriter, req *http.Request) {
	type createChirpParams struct {
		Body   string			`json:"body"`
		UserID uuid.UUID	`json:"user_id"`
	}

	type createChirpResult Chirp 

	decoder := json.NewDecoder(req.Body)
	params := createChirpParams{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanMsg := noNaughtyWords(params.Body)

	if params.UserID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UUID", fmt.Errorf("uuid is not valid"))
	}

	createdChirp, err := a.DB.CreateChirp(req.Context(), database.CreateChirpParams{
		Body:		cleanMsg,
		UserID:	params.UserID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create chirp", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, createChirpResult{
		ID:					createdChirp.ID,
		CreatedAt:	createdChirp.CreatedAt,
		UpdatedAt:	createdChirp.UpdatedAt,
		Body: 			createdChirp.Body,
		UserID:			createdChirp.UserID,
	})
}

func noNaughtyWords(msg string) string {
	const blocked = "****"
	naughtyWords := map[string]struct{} {
		"kerfuffle": 	{},
 		"sharbert": 	{},
 		"fornax": 		{},
	}
	

	tokens := strings.Fields(msg)
	for i, token := range tokens {
		if _, bad := naughtyWords[strings.ToLower(token)]; bad {
			tokens[i] = blocked
		}
	}
	return strings.Join(tokens, " ") 
}

