package routes

import (
	"kasir-api/provider"
	"net/http"
)

func SetupRoute(mux *http.ServeMux, p *provider.Container) {
	mux.HandleFunc("GET /categories", p.CategoryController.CategoryIndex)
	mux.HandleFunc("POST /categories", p.CategoryController.CategoryStore)
	mux.HandleFunc("PUT /categories/{id}", p.CategoryController.CategoryEdit)
	mux.HandleFunc("GET /categories/{id}", p.CategoryController.CategoryFindId)
	mux.HandleFunc("DELETE /categories/{id}", p.CategoryController.CategoryRemove)

	mux.HandleFunc("GET /checkhealth", p.MainController.MainIndex)
}
