package controller

import (
	"log"
	"net/http"
)

type UserController interface {
	BaseController
	Login(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type userController struct {
}

func NewUserController() UserController {
	return &userController{}
}

func (uc userController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("User Request Handler")
}

func (uc userController) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login")
}

func (uc userController) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Create User")
}

func (uc userController) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Update User")
}
