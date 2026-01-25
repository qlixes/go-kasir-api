package main

import (
	"kasir-api/configs"
	"net/http"
)

func main() {

	configs.LoadApp()

	db := configs.LoadDB()

	router := http.NewServeMux()

	server := http.Server{}
}
