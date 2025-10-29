package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/joseph-m-valdez/chirpy/internal/database"
)

func (a *API) HandlerGetChirps(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	authorStr := q.Get("author_id")

	var (
		chirps []database.Chirp
		err    error
	)

	if authorStr != "" {
		authorID, parseErr := uuid.Parse(authorStr)
		if parseErr != nil {
			respondWithError(w, http.StatusBadRequest, "invalid author_id", parseErr)
			return
		}
		chirps, err = a.DB.GetChirpsByUserID(req.Context(), authorID)
	} else {
		chirps, err = a.DB.GetChirps(req.Context())
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error getting chirps", err)
		return
	}

	out := make([]Chirp, 0, len(chirps))
	for _, c := range chirps {
		out = append(out, Chirp{
			ID:        c.ID,
			UserID:    c.UserID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}

	respondWithJSON(w, http.StatusOK, out)
}


func (a *API) HandlerGetChirp(w http.ResponseWriter, req *http.Request) {
	chirpID, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "not a valid uuid", err)
		return
	}

	chirp, err := a.DB.GetChirp(req.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "chirp not found", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
