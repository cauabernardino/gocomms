package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"web/src/config"
	"web/src/controllers"
	"web/src/models"
	"web/src/responses"
	"web/src/utils"

	"github.com/gorilla/mux"
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

// ProfilePage handles the loading of an user profile page.
// It differentiates if you are the logged user or not.
func ProfilePage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	cookie, _ := utils.CheckCookie(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == loggedUserID {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	user, err := controllers.GetAllUserInformation(userID, r)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusInternalServerError,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	RenderTemplate(w, "user.html", struct {
		User         models.User
		LoggedUserID uint64
	}{
		User:         user,
		LoggedUserID: loggedUserID,
	})
}

// LoggedUserProfilePage handles the loading of logged userprofile page.
func LoggedUserProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := utils.CheckCookie(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := controllers.GetAllUserInformation(loggedUserID, r)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusInternalServerError,
			responses.APIError{Error: err.Error()},
		)
		return
	}
	RenderTemplate(w, "profile.html", user)
}
