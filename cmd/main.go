package main

import (
	"log"
	"net/http"
)

func main() {
	// Create API
	api := NewAPI()

	// Create server
	mux := http.NewServeMux()

	// Define routes
	mapRoutes(api, mux)

	// Listen
	log.Fatal(http.ListenAndServe(":8080", mux))
}
