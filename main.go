package main

import (
	"kasir-api/configs"
	"kasir-api/provider"
	"kasir-api/routes"
	"net/http"
)

func main() {
	configs.LoadApp()

	provider := provider.NewContainer()
	mux := http.NewServeMux()

	routes.SetupRoute(mux, provider)

	http.ListenAndServe(":8000", mux)
}
