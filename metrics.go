package main

import (
	"fmt"
	"net/http"
)

// !middleware means first our custom login will happen in middle then the handler will be called
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
