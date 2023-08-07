package main

import (
	"github.com/gorilla/mux"
)

func mapRoutes(api *API, r *mux.Router) {

	// Users
	r.HandleFunc("/login", api.UserController.Login).Methods("POST")
	r.HandleFunc("/users", api.UserController.Create).Methods("POST")
	r.HandleFunc("/users", api.UserController.Update).Methods("PUT")

}
