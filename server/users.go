package server

import (
	"github.com/rebay1982/gostack/db"
	"github.com/rebay1982/gostack/models"
)

func CreateUser(user *models.User) error {

	userDb := user.ToDb()
	err := db.InsertUser(&userDb)

	if err != nil {
		return err
	}
	user.ID = userDb.ID
	return nil
}

func GetUserById(id int) (*models.User, error) {

	var user *models.User = nil
	userDb, err := db.GetUserById(id)

	if userDb != nil {
		tmp := userDb.ToJson()
		user = &tmp
	}

	return user, err
}

func DeleteUserById(id int) error {
	err := db.DeleteUserById(id)
	return err
}
