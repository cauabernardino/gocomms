package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents an users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepository handles link between the struct and database
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create inserts an user to the database
func (repository Users) Create(user models.User) (uint64, error) {
	dbStatement, err := repository.db.Prepare(
		"INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, nil
	}
	defer dbStatement.Close()

	dbExecution, err := dbStatement.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, nil
	}

	lastInsertedID, err := dbExecution.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastInsertedID), nil
}
