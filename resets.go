package main

import "net/http"

func (cfg *apiConfig) resethitsHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
}
