package handler

import (
	"net/http"

	"kasir-api/internal/service"
	"kasir-api/internal/util"
)

type MainHandler struct {
	mainService service.MainService
}

func NewMainHandler(mainService service.MainService) *MainHandler {
	return &MainHandler{
		mainService: mainService,
	}
}

func (c *MainHandler) GetMainIndex(w http.ResponseWriter, r *http.Request) {
	util.ResponseSuccess(w, 200, "API service available", nil)
}
