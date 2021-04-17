package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Search returns all users that comply to the given name or username filter
func (repository Users) Search(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername) // returns %nameOrUsername%

	lines, err := repository.db.Query(
		"SELECT id, name, username, email, created_on FROM users WHERE name LIKE ? OR username LIKE ?",
		nameOrUsername, nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedOn,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
