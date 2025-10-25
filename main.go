package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	"github.com/joseph-m-valdez/chirpy/internal/database"
	"github.com/joseph-m-valdez/chirpy/internal/api"
	"github.com/joseph-m-valdez/chirpy/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	jwtSecret := os.Getenv("JWT_SEKRET")

	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)
	apiCfg := &config.APIConfig{
		FileServerHits: atomic.Int32{},
		DB:							dbQueries,
		Platform:				platform,
		JWTSecret:			jwtSecret,
	}
	api := api.New(apiCfg)

	mux := http.NewServeMux()
	fsHandler := api.MiddlewareMetricsInc(
		http.StripPrefix("/app",
		http.FileServer(http.Dir(filepathRoot))),
	)
	mux.Handle("/app/", fsHandler)
	mux.HandleFunc("GET		/api/healthz", api.HandlerHealth)
	
	mux.HandleFunc("POST	/api/refresh", api.HandlerRefreshToken)
	mux.HandleFunc("POST	/api/revoke", api.HandlerRevokeToken)

	mux.HandleFunc("POST	/api/login", api.HandlerLogin)


	mux.HandleFunc("POST	/api/users", api.HandlerCreateUsers)
	mux.HandleFunc("PUT		/api/users", api.HandlerUpdateUsersAuth)

	mux.HandleFunc("POST			/api/chirps", api.HandlerCreateChirps)
	mux.HandleFunc("GET				/api/chirps", api.HandlerGetChirps)
	mux.HandleFunc("GET				/api/chirps/{chirpID}", api.HandlerGetChirp)
	mux.HandleFunc("DELETE		/api/chirps/{chirpID}", api.HandlerDeleteChirp)

	mux.HandleFunc("GET		/admin/metrics", api.HandlerMetrics)
	mux.HandleFunc("POST	/admin/reset", api.HandlerReset)

	server := &http.Server {
		Addr: 		":" + port,
		Handler:	mux,
	}
	
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}

