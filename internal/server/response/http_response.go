package response

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type HttpResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondWithJson(w http.ResponseWriter, message string, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := HttpResponse{
		Message: message,
		Data:    data,
		Status:  statusCode,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("error marshalling client response: %v", err)
		return
	}

}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if strings.Contains(message, "error marshalling client response") && statusCode != http.StatusTooManyRequests {
		statusCode = http.StatusNotFound
		message = "Resource not found, check you entered the right url and credential"
	}
	if strings.Contains(message, "error marshalling client response") && statusCode == http.StatusTooManyRequests {
		message = "Rate limit exceeded"
	}

	response := HttpResponse{
		Error:  message,
		Status: statusCode,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("error marshalling client response: %v", err)
		return
	}
}
