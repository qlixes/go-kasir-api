package controllers

import (
	"encoding/json"
	"net/http"
)

type MainController struct {
}

func NewMainController() *MainController {
	return &MainController{}
}

func (c *MainController) MainIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	format := map[string]string{
		"status":  "Success",
		"message": "API running",
	}

	json.NewEncoder(w).Encode(format)
}
