package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser handles the creation of new user in the platform.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ReturnError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepository := repositories.NewUsersRepository(db)
	user.ID, err = usersRepository.Create(user)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusCreated, user)
}

// SearchUsers handles the searching for all users of the platform.
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching for all users!"))
}

// SearchUser handles the searching for a specific user.
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching for a specific user!"))
}

// UpdateUser handles the edition and/or updating of an user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating an user!"))
}

// DeleteUser handles the deletion of an user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting an user!"))
}
