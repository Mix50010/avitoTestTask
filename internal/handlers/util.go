package handlers

import (
	"encoding/json"
	"net/http"
	"tenderService/internal/models"
)

func WriteError(w http.ResponseWriter, reason string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(models.ErrorResponse{Reason: reason})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
