package util

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(w http.ResponseWriter, httpCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	response := HttpResponse{
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func ResponseFail(w http.ResponseWriter, httpCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	response := HttpResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
