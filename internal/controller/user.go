package controller

import (
	"errors"
	"github.com/gonzamedrano09/join_up/internal/service"
	"log"
	"net/http"
)

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) (UserController, error) {
	if userService == nil {
		return nil, errors.New("userService can not be nil")
	}

	return &userController{
		userService: userService,
	}, nil
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
