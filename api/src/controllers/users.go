package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/utils"
	"encoding/json"
	"errors"
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

	tokenUserID, err := utils.ExtractUserID(r)
	if err != nil {
		responses.ReturnError(w, http.StatusUnauthorized, err)
		return
	}

	// Check for permission
	if userID != tokenUserID {
		err := errors.New("you're not allowed to perform this action")
		responses.ReturnError(w, http.StatusForbidden, err)
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

	tokenUserID, err := utils.ExtractUserID(r)
	if err != nil {
		responses.ReturnError(w, http.StatusUnauthorized, err)
	}

	// Check for permission
	if userID != tokenUserID {
		err := errors.New("you're not allowed to perform this action")
		responses.ReturnError(w, http.StatusForbidden, err)
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

// FollowUser allows an authenticated user to follow other user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := utils.ExtractUserID(r)
	if err != nil {
		responses.ReturnError(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	// Check for same ID
	if followerID == userID {
		err := errors.New("you can't follow yourself")
		responses.ReturnError(w, http.StatusForbidden, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepository := repositories.NewUsersRepository(db)

	if err = usersRepository.Follow(userID, followerID); err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusNoContent, nil)
}

// UnfollowUser allows an authenticated user to unfollow other user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := utils.ExtractUserID(r)
	if err != nil {
		responses.ReturnError(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	// Check for same ID
	if followerID == userID {
		err := errors.New("you can't unfollow yourself")
		responses.ReturnError(w, http.StatusForbidden, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepository := repositories.NewUsersRepository(db)

	if err = usersRepository.Unfollow(userID, followerID); err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusNoContent, nil)
}

// SearchFollowers handles the search for all followers from an user
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
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
	followers, err := userRepository.Followers(userID)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusOK, followers)
}

// SearchFollowing handles the search for who the user is following
func SearchFollowing(w http.ResponseWriter, r *http.Request) {
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
	following, err := userRepository.Following(userID)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusOK, following)
}

// ResetPassword handles the password reset of an user
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := utils.ExtractUserID(r)
	if err != nil {
		responses.ReturnError(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	// Check for permission
	if tokenUserID != userID {
		err := errors.New("you're not allowed to perform this action")
		responses.ReturnError(w, http.StatusForbidden, err)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ReturnError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password
	if err = json.Unmarshal(reqBody, &password); err != nil {
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
	passwordInDB, err := userRepository.SearchPassword(userID)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	// Verifying current password
	if err = utils.VerifyPassword(passwordInDB, password.Current); err != nil {
		err = errors.New("wrong password")
		responses.ReturnError(w, http.StatusUnauthorized, err)
	}

	// Hashing and saving new password
	hashedPassword, err := utils.GenerateHash(password.New)
	if err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	if err = userRepository.UpdatePassword(userID, string(hashedPassword)); err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusNoContent, nil)
}
