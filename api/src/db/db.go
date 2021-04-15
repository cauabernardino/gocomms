package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect opens the connection with the database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBConnectString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
