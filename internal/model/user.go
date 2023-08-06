package model

type User struct {
	Username       string
	Email          string
	HashedPassword string

	Picture   string
	FirstName string
	LastName  string
	Biography string
	Location  Location

	Tags []Tag
}
