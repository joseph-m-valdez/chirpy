package api

import (
	"net/http"

	"github.com/joseph-m-valdez/chirpy/internal/auth"
)

func (a *API) HandlerRevokeToken(w http.ResponseWriter, req *http.Request) {
   bearer, err := auth.GetBearerToken(req.Header)
	 if err != nil {
			respondWithError(w, http.StatusInternalServerError, "no bearer token", err)
		 	return
	 }

		if err := a.DB.RevokeRefreshToken(req.Context(), bearer); err != nil {
			respondWithError(w, http.StatusInternalServerError, "cannot revoke refresh token", err)
		 	return
		}

		w.WriteHeader(http.StatusNoContent)
	
}
