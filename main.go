package main

import (
	"net/http"

	"github.com/mingchengzack/REST-API/api"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get"}`))
}

func main() {
	api := api.NewAPI()
	api.Get("/", get)
	api.Run(":8080")
}
