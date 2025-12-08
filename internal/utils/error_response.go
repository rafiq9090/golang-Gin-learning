package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details"`
}

func JSONError(w http.ResponseWriter, message string, statusCode int, details map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ErrorResponse{
		Error:   message,
		Details: details,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to send JSON error response: %v", err)
	}
}

func JSONSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to send JSON success response: %v", err)
	}
}
