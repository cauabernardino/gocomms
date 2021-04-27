package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)
	registeredUser, err := userRepository.SearchEmail(user.Email)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
	}

	if err = utils.VerifyPassword(registeredUser.Password, user.Password); err != nil {
		responses.ReturnError(w, http.StatusUnauthorized, err)
		return
	}

	token, err := utils.CreateToken(registeredUser.ID)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(registeredUser.ID, 10)
	authData := models.AuthData{ID: userID, Token: token}

	responses.ReturnJSON(w, http.StatusOK, authData)
}
