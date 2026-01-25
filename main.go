package main

import (
	"kasir-api/configs"
	"net/http"
)

func main() {

	configs.LoadApp()

	err := http.ListenAndServe("8000", nil)

	if err != nil {
		panic("Error run server")
	}
}
