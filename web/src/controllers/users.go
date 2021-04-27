package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"web/src/responses"
)

// CreateUser handles the call to API for registering an user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"username": r.FormValue("username"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	response, err := http.Post(
		"http://localhost:5000/users",
		"application/json",
		bytes.NewBuffer(user),
	)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusInternalServerError,
			responses.APIError{Error: err.Error()},
		)
		return

	}
	defer response.Body.Close()

	// Check if the status code returned from API is an error
	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	responses.ReturnJSON(w, response.StatusCode, nil)
}
