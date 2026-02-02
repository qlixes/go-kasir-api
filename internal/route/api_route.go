package routes

import (
	"net/http"

	"kasir-api/internal/provider"
)

func SetupRoute(mux *http.ServeMux, p *provider.HandlerProvider) {
	// mux.HandleFunc("GET /products", p.ProductController.ProductIndex)
	// mux.HandleFunc("POST /products", p.ProductController.ProductStore)
	// mux.HandleFunc("PUT /products/{id}", p.ProductController.ProductEdit)
	// mux.HandleFunc("GET /products/{id}", p.ProductController.ProductFindId)
	// mux.HandleFunc("DELETE /products/{id}", p.ProductController.ProductRemove)

	mux.HandleFunc("GET /categories", p.CategoryHandler.GetCategoryIndex)
	mux.HandleFunc("POST /categories", p.CategoryHandler.PostCategoryStore)
	mux.HandleFunc("PUT /categories/{id}", p.CategoryHandler.PutCategoryId)
	mux.HandleFunc("GET /categories/{id}", p.CategoryHandler.GetCategoryId)
	mux.HandleFunc("DELETE /categories/{id}", p.CategoryHandler.DeleteCategoryId)

	mux.HandleFunc("GET /checkhealth", p.MainHandler.MainIndex)
}
