package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.Prepare("SignUp"); err != nil {
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
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepository := repositories.NewUsersRepository(db)
	users, err := usersRepository.Search(nameOrUsername)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusOK, users)
}

// SearchUser handles the searching for a specific user.
func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)
	user, err := userRepository.SearchID(userID)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusOK, user)
}

// UpdateUser handles the edition and/or updating of an user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

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

	if err = user.Prepare("Update"); err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)
	if err := userRepository.Update(userID, user); err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusNoContent, nil)
}

// DeleteUser handles the deletion of an user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)
	if err := userRepository.Delete(userID); err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusNoContent, nil)
}
