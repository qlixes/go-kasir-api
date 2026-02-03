package util

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(w http.ResponseWriter, httpCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	response := HttpResponse{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func ResponseFail(w http.ResponseWriter, httpCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	response := HttpResponse{
		Code:    httpCode,
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
