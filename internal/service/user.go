package service

import (
	"context"
	"errors"
	"github.com/gonzamedrano09/join_up/internal/model"
	"github.com/gonzamedrano09/join_up/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, userCreateInput *model.UserCreateInput) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) (UserService, error) {
	if userRepository == nil {
		return nil, errors.New("userRepository can not be nil")
	}

	return userService{
		userRepository: userRepository,
	}, nil
}

func (us userService) Create(ctx context.Context, userCreateInput *model.UserCreateInput) (*model.User, error) {
	if err := userCreateInput.IsValid(); err != nil {
		return nil, err
	}

	user, err := model.NewUser(userCreateInput)
	if err != nil {
		return nil, err
	}

	if err = us.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
