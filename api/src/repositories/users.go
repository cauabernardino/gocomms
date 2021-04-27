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
		return 0, err
	}
	defer dbStatement.Close()

	dbExecution, err := dbStatement.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := dbExecution.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

// Search returns all users that comply to the given name or username filter
func (repository Users) Search(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername) // returns %nameOrUsername%

	rows, err := repository.db.Query(
		"SELECT id, name, username, email, created_on FROM users WHERE name LIKE ? OR username LIKE ?",
		nameOrUsername, nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
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

// SearchID returns the user information of a given ID if it exists in database
func (repository Users) SearchID(ID uint64) (models.User, error) {
	row, err := repository.db.Query(
		"SELECT id, name, username, email, created_on FROM users WHERE id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	for row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedOn,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update handles the update of user information in database
func (repository Users) Update(ID uint64, user models.User) error {
	dbStatement, err := repository.db.Prepare(
		"UPDATE users SET name = ?, username = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer dbStatement.Close()

	if _, err = dbStatement.Exec(user.Name, user.Username, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Delete handles the deletion of an user from the database
func (repository Users) Delete(ID uint64) error {
	dbStatement, err := repository.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer dbStatement.Close()

	if _, err = dbStatement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// SearchEmail returns an user information if the given email exists in database
func (repository Users) SearchEmail(email string) (models.User, error) {
	row, err := repository.db.Query(
		"SELECT id, password FROM users WHERE email = ?",
		email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Follow inserts the follow made by an user into the database, if possible
func (repository Users) Follow(userID, followerID uint64) error {
	dbStatement, err := repository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer dbStatement.Close()

	if _, err = dbStatement.Exec(userID, followerID); err != nil {
		return err
	}
	return nil
}

// Unfollow deletes a follow made by an user from the database, if possible
func (repository Users) Unfollow(userID, followerID uint64) error {
	dbStatement, err := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? and follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer dbStatement.Close()

	if _, err = dbStatement.Exec(userID, followerID); err != nil {
		return err
	}
	return nil
}

// Followers returns a list with the followers of a given userID
func (repository Users) Followers(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
	SELECT u.id, u.name, u.username, u.email, u.created_on
	FROM users u INNER JOIN followers f ON u.id = f.follower_id
	WHERE f.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
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

// Following returns a list with the users tha the given userID is following
func (repository Users) Following(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
	SELECT u.id, u.name, u.username, u.email, u.created_on
	FROM users u INNER JOIN followers f ON u.id = f.user_id
	WHERE f.follower_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
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

// SearchPassword returns the password of an user by userID
func (repository Users) SearchPassword(userID uint64) (string, error) {
	row, err := repository.db.Query(
		"SELECT password FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// UpdatePassword updates the password of an user in database
func (repository Users) UpdatePassword(userID uint64, hashedPassword string) error {
	dbStatement, err := repository.db.Prepare(
		"UPDATE users SET password = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer dbStatement.Close()

	if _, err = dbStatement.Exec(hashedPassword, userID); err != nil {
		return err
	}

	return nil
}
