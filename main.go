package main

import (
	"fmt"
	"log"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	const filepathRoot = "."
	const port = "8080"
	cfg := &apiConfig{
		fileserverHits: 0,
	}
	mux := http.NewServeMux()

	handler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))

	mux.Handle("/app/*", cfg.middlewareMetricsInc(handler))

	mux.HandleFunc("/healthz", readinessHandler)
	mux.HandleFunc("/metrics", cfg.hitsHandler)
	mux.HandleFunc("/reset", cfg.resethitsHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())

}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		//!chain of handler so serveHTTP hands off control to the "next" handler
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) hitsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hits: %v", cfg.fileserverHits)
}
func (cfg *apiConfig) resethitsHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
}
