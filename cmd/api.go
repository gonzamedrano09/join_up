package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/gonzamedrano09/join_up/internal/config"
	"github.com/gonzamedrano09/join_up/internal/controller"
	"github.com/gonzamedrano09/join_up/internal/repository"
	"github.com/gonzamedrano09/join_up/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	usersCollection = "users"
)

type API struct {
	clients struct {
		dbConnection *mongo.Client
	}

	repositories struct {
		userRepository repository.UserRepository
	}

	services struct {
		userService service.UserService
	}

	controllers struct {
		userController controller.UserController
	}
}

func NewAPI() (*API, error) {
	// Init API
	api := new(API)

	// Create clients
	dbConnection, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/",
			config.Config.DatabaseUser,
			config.Config.DatabasePassword,
			config.Config.DatabaseHost,
			config.Config.DatabasePort,
		)),
	)
	if err != nil {
		return nil, errors.New("error creating dbConnection")
	}
	api.clients.dbConnection = dbConnection

	// Create repositories
	userRepository, err := repository.NewUserRepository(dbConnection, config.Config.DatabaseName, usersCollection)
	if err != nil {
		return nil, errors.New("error creating userRepository")
	}
	api.repositories.userRepository = userRepository

	// Create services
	userService, err := service.NewUserService(userRepository)
	if err != nil {
		return nil, errors.New("error creating userService")
	}
	api.services.userService = userService

	// Create controllers
	userController, err := controller.NewUserController(userService)
	if err != nil {
		return nil, errors.New("error creating userController")
	}
	api.controllers.userController = userController

	// Return pointer to API struct
	return api, nil
}
