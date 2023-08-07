package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create API
	api := NewAPI()

	// Create router
	r := mux.NewRouter()

	// Define routes
	mapRoutes(api, r)

	// Listen
	log.Fatal(http.ListenAndServe(":8080", r))
}
