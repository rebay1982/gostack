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

	userDb, err := db.GetUserById(id)

	if err != nil {

		return nil, err
	}

	if userDb == nil {
		return nil, nil
	}

	user := userDb.ToJson()
	return &user, nil
}
