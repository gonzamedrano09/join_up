package main

import (
	"net/http"
)

func mapRoutes(api *API, m *http.ServeMux) {

	m.HandleFunc("/users", api.UserController.HandleRequest)

}
