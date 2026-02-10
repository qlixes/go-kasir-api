package routes

import (
	"net/http"

	"kasir-api/internal/provider"
)

func SetupRoute(mux *http.ServeMux, p *provider.HandlerProvider) {
	mux.HandleFunc("GET /products", p.ProductHandler.GetProductIndex)
	mux.HandleFunc("POST /products", p.ProductHandler.PostProductStore)
	mux.HandleFunc("PUT /products/{id}", p.ProductHandler.PutProductId)
	mux.HandleFunc("GET /products/{id}", p.ProductHandler.GetProductId)
	mux.HandleFunc("DELETE /products/{id}", p.ProductHandler.DeleteProductId)

	mux.HandleFunc("GET /categories", p.CategoryHandler.GetCategoryIndex)
	mux.HandleFunc("POST /categories", p.CategoryHandler.PostCategoryStore)
	mux.HandleFunc("PUT /categories/{id}", p.CategoryHandler.PutCategoryId)
	mux.HandleFunc("GET /categories/{id}", p.CategoryHandler.GetCategoryId)
	mux.HandleFunc("DELETE /categories/{id}", p.CategoryHandler.DeleteCategoryId)

	mux.HandleFunc("GET /checkhealth", p.MainHandler.GetMainIndex)

	// mux.HandleFunc("GET /report", p.ReportHandler.GetMainIndex)
	// mux.HandleFunc("POST /checkout", p.CheckoutHandler.GetMainIndex)
}
