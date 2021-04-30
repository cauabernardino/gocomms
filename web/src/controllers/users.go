package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"web/src/config"
	"web/src/responses"
	"web/src/utils"
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

	url := fmt.Sprintf("%s/users", config.APIURL)
	response, err := http.Post(
		url,
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
		responses.HandleAPIStatusCodeError(w, response)
		return
	}

	responses.ReturnJSON(w, response.StatusCode, nil)
}

// UserAuthenticatedRequest handles the token insertion into request
func UserAuthenticatedRequest(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	// Creates the request
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	// Get cookie from browser and add it to request header
	cookie, _ := utils.CheckCookie(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	// Execution of request and catching response
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
