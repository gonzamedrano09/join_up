package model

type User struct {
	Email          string
	HashedPassword string
	FirstName      string
	LastName       string
	Biography      string
	Tags           []Tag
	// Picture
	// Location
}
