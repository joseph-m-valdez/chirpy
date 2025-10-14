package api

import (
	"net/http"
	"github.com/joseph-m-valdez/chirpy/internal/config"
)

type API struct {
	*config.APIConfig
}

func New(c *config.APIConfig) *API {
	return &API{APIConfig: c}
}

func (a *API) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		a.FileServerHits.Add(1)
		next.ServeHTTP(w, req)
	})
}
