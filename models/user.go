package models

// UserDb User model for use against the database.
type UserDb struct {
	ID        int
	FirstName string
	LastName  string
}

// UserDbList List of user models for use against the database.
type UserDbList struct {
	UsersDb []UserDb
}

// User User model for use against the HTTP API.
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// UserList User list for use against the HTTP API.
type UserList struct {
	Users []User `json:"users"`
}
