package db

import (
	"database/sql"
	"github.com/rebay1982/gostack/models"

	_ "github.com/lib/pq"
)

type database struct {
	cnn *sql.Db
}

func initialize(isername, password, database string) (database, error) {

	db := database{}
	connStr := "host=database port=5432 user=gostack dbname=gostack sslmode=off"

	cnn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.cnn = cnn

	return db
}

// GetUserById returns a UserDb by associated with the ID passed in paramter.
func (db database) GetUserById(id int) (*UserDb, error) {
	user := &models.UserDb{}

	query := `SELECT * FROM users WHERE id = $1;`

	row := db.cnn.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	return user, err
}

// InsertUser takes a UserDb as parameter and inserts it in the data base.
//   This method returns an error if, for whatever reason, it was not possible
//   to insert the User.
func (db database) InsertUser(user *models.UserDb) error {
	var id int
	query := `INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING id`

	err := db.cnn.QueryRow(query, user.FirstName, user.LastName).Scan(&id)

	if err != nil {
		return err
	}

	user.ID = id
	return nil
}
