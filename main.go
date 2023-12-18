package main

import (
	"github.com/Harsharma2026/car-garage-api/handlers"
	"net/http"
)

var router *mux.Router = handlers.NewRouter()

func main() {
	http.ListenAndServe(":8080", router)
}
