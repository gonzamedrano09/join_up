package main

import (
	"github.com/gorilla/mux"
)

func mapRoutes(api *API, r *mux.Router) {

	// Users
	r.HandleFunc("/login", api.controllers.userController.Login).Methods("POST")
	r.HandleFunc("/users", api.controllers.userController.Create).Methods("POST")
	r.HandleFunc("/users", api.controllers.userController.Update).Methods("PUT")

}
