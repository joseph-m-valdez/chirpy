package api

import (
	"fmt"
	"net/http"
)

func (a *API) HandlerReset(w http.ResponseWriter, req *http.Request) {
	platform := a.Platform

	if platform != "dev" {
		err := fmt.Errorf("forbidden")
		respondWithError(w, http.StatusForbidden, "forbidden", err)
		return
	}

	err := a.DB.DeleteUsers(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't delete users", err)
		return
	}

	respondWithJSON(w, http.StatusOK, "all users deleted")
}
