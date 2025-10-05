package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.FileServer(http.Dir(filepathRoot)))
	mux.HandleFunc("/healthz", healthHandler)

	server := &http.Server {
		Addr: 		":" + port,
		Handler:	mux,
	}
	
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
