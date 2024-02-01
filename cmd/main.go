package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main/internal/config"
)

func main() {
	// Load config
	if err := config.LoadConfig(); err != nil {
		return
	}

	// Create API
	api := NewAPI()

	// Create router
	r := mux.NewRouter()

	// Define routes
	mapRoutes(api, r)

	// Listen
	log.Fatal(http.ListenAndServe(":8080", r))
}
