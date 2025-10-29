package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/joseph-m-valdez/chirpy/internal/auth"
	"github.com/joseph-m-valdez/chirpy/internal/database"
)

func (a *API) HandlerPolkaWebHooks(w http.ResponseWriter, req *http.Request) {
	type requestParams struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	var reqParams requestParams

	polkaApiKey, err := auth.GetAPIKey(req.Header)

	if polkaApiKey != a.PolkaKey {
		respondWithError(w, http.StatusUnauthorized, "Invalid api key", err)
		return
	}

	if err := json.NewDecoder(req.Body).Decode(&reqParams); err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
		return
	}

	if strings.ToLower(reqParams.Event) != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, nil)
	}

	userID, err := uuid.Parse(reqParams.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error processing request, invalid UUID", err)
		return
	}

	_, err = a.DB.UpdateUserMembership(req.Context(), database.UpdateUserMembershipParams{
		ID:          userID,
		IsChirpyRed: true,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Failed to update membership", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil )
}
