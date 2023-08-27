package model

import (
	"encoding/json"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Inputs

type UserCreateInput struct {
	Username string
	Password string

	FirstName string
	LastName  string
}

func (u *UserCreateInput) IsValid() bool {
	return true
}

type UserUpdateInput struct {
	FirstName string
	LastName  string
	Biography string
}

func (u *UserUpdateInput) IsValid() bool {
	return true
}

type UserUpdateLocationsInput struct {
	Locations []Location
}

func (u *UserUpdateLocationsInput) IsValid() bool {
	return true
}

type UserUpdateTagsInput struct {
	Tags []Tag
}

func (u *UserUpdateTagsInput) IsValid() bool {
	return true
}

// Model

type User struct {
	Username       string
	Email          string
	HashedPassword string

	Picture   string
	FirstName string
	LastName  string
	Biography string

	Locations []Location
	Tags      []Tag
}

func NewUser(ui *UserCreateInput) (*User, error) {
	userBytes, err := json.Marshal(ui)
	if err != nil {
		return nil, errors.New("error creating user")
	}

	user := User{}
	if err = json.Unmarshal(userBytes, &user); err != nil {
		return nil, errors.New("error creating user")
	}

	if err = user.SetPassword(ui.Password); err != nil {
		return nil, errors.New("error creating user")
	}

	return &user, nil
}

func (u *User) SetPassword(password string) error {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return errors.New("error setting password user")
	}

	u.HashedPassword = string(passwordBytes)

	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}
