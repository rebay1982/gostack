package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

type database struct {
	cnn *sql.DB
}

//var once sync.Once
var (
	lock = &sync.Mutex{}
	db   = database{}
)

func getDb() (database, error) {
	lock.Lock()
	defer lock.Unlock()

	if db.cnn == nil {
		connStr := "host=database port=5432 user=gostack dbname=gostack password=gostack_password sslmode=disable"

		cnn, err := sql.Open("postgres", connStr)
		if err != nil {
			return db, err
		}

		db.cnn = cnn
	}

	return db, nil
}
