package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"web/src/responses"
)

// UserLogin uses email and password from the user to perform login
func UserLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
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
		"http://localhost:5000/login",
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

	token, _ := ioutil.ReadAll(response.Body)
	fmt.Println(response.StatusCode, string(token))
}
