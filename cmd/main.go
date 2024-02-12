package main

import (
	"log"
	"net/http"

	"github.com/gonzamedrano09/join_up/internal/config"
	"github.com/gorilla/mux"
)

func main() {
	// Load config
	if err := config.LoadConfig(); err != nil {
		panic("error loading config")
	}

	// Create API
	api, err := NewAPI()
	if err != nil {
		panic(err)
	}

	// Create router
	r := mux.NewRouter()

	// Define routes
	mapRoutes(api, r)

	// Listen
	log.Fatal(http.ListenAndServe(":8080", r))
}
