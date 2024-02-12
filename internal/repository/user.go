package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/gonzamedrano09/join_up/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
}

type userRepository struct {
	dbConnection   *mongo.Client
	databaseName   string
	collectionName string
}

func NewUserRepository(dbConnection *mongo.Client, databaseName string, collectionName string) (UserRepository, error) {
	if dbConnection == nil {
		return nil, errors.New("dbConnection can not be nil")
	}

	if databaseName == "" {
		return nil, errors.New("databaseName can not be empty")
	}

	if collectionName == "" {
		return nil, errors.New("collectionName can not be empty")
	}

	return userRepository{
		dbConnection:   dbConnection,
		databaseName:   databaseName,
		collectionName: collectionName,
	}, nil
}

func (ur userRepository) Create(ctx context.Context, user *model.User) error {
	collection := ur.dbConnection.Database(ur.databaseName).Collection(ur.collectionName)

	if _, err := collection.InsertOne(ctx, user); err != nil {
		return errors.New(fmt.Sprintf("error creating user: %s", err))
	}

	return nil
}
