package main

import "main/internal/controller"

type API struct {
	UserController controller.UserController
}

func NewAPI() *API {
	// Define controllers
	userController := controller.NewUserController()

	// Return pointer to API struct
	return &API{UserController: userController}
}
