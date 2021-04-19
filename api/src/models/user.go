package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents the structure of an user data.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
}

// Prepare calls the validation and formatting functions for
// the User struct
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
	return nil
}

// validate the User struct
func (user *User) validate(step string) error {

	if user.Name == "" {
		return errors.New("field name is required")
	}

	if user.Username == "" {
		return errors.New("field username is required")
	}

	if user.Email == "" {
		return errors.New("field email is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("inserted email is not valid")
	}

	if step == "SignUp" && user.Password == "" {
		return errors.New("field password is required")
	}

	return nil
}

// formar formats the values for the User struct
func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
}
