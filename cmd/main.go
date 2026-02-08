package main

import (
	"fmt"
	"log"
	"net/http"

	"kasir-api/internal/infrastructure"
	"kasir-api/internal/provider"
	routes "kasir-api/internal/route"
)

func main() {
	config := infrastructure.NewConfig()
	appPort := fmt.Sprintf(":%d", config.AppConfig.Port)

	db, err := infrastructure.NewPgsql(config)
	if err != nil {
		log.Fatalln("Failed connect database !")
	}
	defer db.Close()

	// DI
	repo := provider.NewRepository(db)
	svc := provider.NewService(repo)
	handlers := provider.NewHandler(svc)

	mux := http.NewServeMux()
	routes.SetupRoute(mux, handlers)

	if err := http.ListenAndServe(appPort, mux); err != nil {
		log.Fatalln("Failed run server !")
	}
}
