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

// ToJson Takes a UserDb and converts it to a User
func (u *UserDb) ToJson() User {
	user := User{}
	user.ID = u.ID
	user.FirstName = u.FirstName
	user.LastName = u.LastName

	return user
}

// ToDb Takes a User and converts it to a UserDb
func (u *User) ToDb() UserDb {
	userDb := UserDb{}
	userDb.ID = u.ID
	userDb.FirstName = u.FirstName
	userDb.LastName = u.LastName

	return userDb
}
