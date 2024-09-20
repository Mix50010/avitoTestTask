package handlers

import (
	"log"
	"net/http"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing ping request...")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("ok"))
}
