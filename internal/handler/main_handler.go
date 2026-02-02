package handler

import (
	"encoding/json"
	"kasir-api/internal/service"
	"net/http"
)

type MainHandler struct {
	mainService service.MainService
}

func NewMainHandler(mainService service.MainService) *MainHandler {
	return &MainHandler{
		mainService: mainService,
	}
}

func (c *MainHandler) MainIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	format := map[string]string{
		"status":  "Success",
		"message": "API running",
	}

	json.NewEncoder(w).Encode(format)
}
