package routes

import (
	"kasir-api/controllers"
	"net/http"
)

func SetupRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /categories", controllers.CategoryIndex)
	mux.HandleFunc("POST /categories", container.CategoryStore)
	mux.HandleFunc("PUT /categories/{id}", container.CategoryEdit)
	mux.HandleFunc("GET /categories/{id}", container.CategoryFind)
	mux.HandleFunc("DELETE /categories/{id}", container.CategoryRemove)
}
