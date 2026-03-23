package utils

import (
	"encoding/json"
	"net/http"

	"github.com/aprimr/chautari/models"
)

func SendSuccess(w http.ResponseWriter, message string, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Success: false,
		Message: message,
	})
}
