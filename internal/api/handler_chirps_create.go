package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/joseph-m-valdez/chirpy/internal/auth"
	"github.com/joseph-m-valdez/chirpy/internal/database"
)

func (a *API) HandlerCreateChirps(w http.ResponseWriter, req *http.Request) {
    type createChirpParams struct {
        Body string `json:"body"`
    }
    type createChirpResult Chirp

    var params createChirpParams
    if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
        respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
        return
    }

    bearer, err := auth.GetBearerToken(req.Header)
    if err != nil {
        respondWithError(w, http.StatusUnauthorized, "missing or invalid auth", err)
        return
    }

    userID, err := auth.ValidateJWT(bearer, a.JWTSecret)
    if err != nil {
        respondWithError(w, http.StatusUnauthorized, "invalid token", err)
        return
    }

    const maxChirpLength = 140
    if len(params.Body) > maxChirpLength {
        respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
        return
    }

    cleanMsg := noNaughtyWords(params.Body)

    createdChirp, err := a.DB.CreateChirp(req.Context(), database.CreateChirpParams{
        Body:   cleanMsg,
        UserID: userID,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "couldn't create chirp", err)
        return
    }

    respondWithJSON(w, http.StatusCreated, createChirpResult{
        ID:        createdChirp.ID,
        CreatedAt: createdChirp.CreatedAt,
        UpdatedAt: createdChirp.UpdatedAt,
        Body:      createdChirp.Body,
        UserID:    createdChirp.UserID,
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

