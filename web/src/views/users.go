package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"web/src/config"
	"web/src/controllers"
	"web/src/models"
	"web/src/responses"
)

// UsersPage handles the loading of queried users
func UsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	url := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrUsername)
	response, err := controllers.UserAuthenticatedRequest(
		r,
		http.MethodGet,
		url,
		nil,
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

	if response.StatusCode >= 400 {
		responses.HandleAPIStatusCodeError(w, response)
		return
	}

	var users []models.User
	if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.ReturnJSON(
			w,
			http.StatusUnprocessableEntity,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	RenderTemplate(w, "users.html", users)
}
