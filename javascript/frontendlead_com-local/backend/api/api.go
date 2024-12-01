package api

import (
	"log"
	"net/http"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	log.Printf("url path: %s", r.URL.Path)
	switch r.URL.Path {
	case "/api/run":
		handleRun(w, r)
	case "/api/questions":
		handleGetDetail(w, r)
	default:
		notFound(w, r)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}
