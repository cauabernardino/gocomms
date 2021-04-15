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
func (u Users) Create(user models.User) (uint64, error) {

	return 0, nil
}
