package api

import (
	"net/http"
	"sort"
	"strings"

	"github.com/google/uuid"
	"github.com/joseph-m-valdez/chirpy/internal/database"
)

func (a *API) HandlerGetChirps(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	authorStr := q.Get("author_id")
	sortParam := strings.ToLower(q.Get("sort")) // "", "asc", or "desc"

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

	// Only sort in-memory if sort is explicitly provided.
	// Ideally we would just have specific queries for asc and desc
	// but this is fine for now
	if sortParam != "" {
		if sortParam != "asc" && sortParam != "desc" {
			respondWithError(w, http.StatusBadRequest, "invalid sort (use asc or desc)", nil)
			return
		}
		asc := sortParam == "asc"
		sort.Slice(chirps, func(i, j int) bool {
			ci, cj := chirps[i], chirps[j]
			if !ci.CreatedAt.Equal(cj.CreatedAt) {
				if asc {
					return ci.CreatedAt.Before(cj.CreatedAt)
				}
				return ci.CreatedAt.After(cj.CreatedAt)
			}
			return ci.ID.String() < cj.ID.String()
		})
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
