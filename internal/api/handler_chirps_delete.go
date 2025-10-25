package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/joseph-m-valdez/chirpy/internal/auth"
	"github.com/joseph-m-valdez/chirpy/internal/database"
)

func (a *API) HandlerDeleteChirp(w http.ResponseWriter, req *http.Request) {
    chirpID, err := uuid.Parse(req.PathValue("chirpID"))
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "no chirp id provided", err)
        return
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

    chirp, err := a.DB.GetChirp(req.Context(), chirpID)
    if err != nil {
        respondWithError(w, http.StatusNotFound, "chirp not found", err)
        return
    }

    if chirp.UserID != userID {
        respondWithError(w, http.StatusForbidden, "cannot delete chirp", nil)
        return
    }

    if err := a.DB.DeleteChirp(req.Context(), database.DeleteChirpParams{
        ID: chirpID, UserID: userID,
    }); err != nil {
        respondWithError(w, http.StatusInternalServerError, "could not delete chirp", err)
        return
    }

    respondWithJSON(w, http.StatusNoContent, nil)
}
