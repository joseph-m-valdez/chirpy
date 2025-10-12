package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	platform := cfg.platform

	if platform != "dev" {
		err := fmt.Errorf("forbidden")
		respondWithError(w, http.StatusForbidden, "forbidden", err)
		return
	}

	err := cfg.db.DeleteUsers(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't delete users", err)
		return
	}

	respondWithJSON(w, http.StatusOK, "all users deleted")
}
