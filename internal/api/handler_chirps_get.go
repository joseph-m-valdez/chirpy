package api

import (
	"net/http"

	"github.com/google/uuid"
)

func (a *API) HandlerGetChirps(w http.ResponseWriter, req *http.Request) {
	chirps, err := a.DB.GetChirps(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError," error getting chirps", err)
	}

	out := make([]Chirp, len(chirps))
	for i, c := range chirps {
		out[i] = Chirp{
			ID:					c.ID,
			UserID:			c.UserID,
			Body:				c.Body,
			CreatedAt:	c.CreatedAt,
			UpdatedAt:	c.UpdatedAt,
		}
	}

	respondWithJSON(w, http.StatusOK, out)

}

func (a *API) HandlerGetChirp(w http.ResponseWriter, req *http.Request) {
	chirpID, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "not a valid uuid", err)
		return
	}

	chirp, err := a.DB.GetChirp(req.Context(), chirpID )

	if err != nil {
		respondWithError(w, http.StatusNotFound, "chirp not found", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID: 				chirp.ID,
		CreatedAt:	chirp.CreatedAt,
		UpdatedAt:	chirp.UpdatedAt,
		Body:				chirp.Body,
		UserID:			chirp.UserID,
	})
}
