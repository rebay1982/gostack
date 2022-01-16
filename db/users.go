package db

import (
	"database/sql"
	"github.com/rebay1982/gostack/models"
)

// GetUserById returns a UserDb by associated with the ID passed in paramter.
func GetUserById(id int) (*models.UserDb, error) {
	userDb := &models.UserDb{}
	query := `SELECT * FROM users WHERE id = $1;`

	db, err := getDb()
	if err != nil {
		return nil, err
	}

	row := db.cnn.QueryRow(query, id)

	err = row.Scan(&userDb.ID, &userDb.FirstName, &userDb.LastName)

	// Shouldn't consider no rows as an error.
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return userDb, nil
}

// InsertUser takes a UserDb as parameter and inserts it in the data base.
//   This method returns an error if, for whatever reason, it was not possible
//   to insert the User.
func InsertUser(user *models.UserDb) error {
	var id int
	query := `INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING id;`

	db, err := getDb()
	if err != nil {
		return err
	}

	err = db.cnn.QueryRow(query, user.FirstName, user.LastName).Scan(&id)

	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

func DeleteUserById(id int) error {
	query := `DELETE FROM users WHERE id = $1;`

	db, err := getDb()
	if err != nil {
		return err
	}

	_, err = db.cnn.Exec(query, id)
	switch err {
	case sql.ErrNoRows:
		return nil

	default:
		return err // If no errors were produced, we're returning nil.
	}
}
