// Package config: api configuration
package config

import (
	"sync/atomic"

	"github.com/joseph-m-valdez/chirpy/internal/database"
)

type APIConfig struct {
	FileServerHits 	atomic.Int32
	DB			 			 	*database.Queries
	Platform				string	
	JWTSecret				string
}

