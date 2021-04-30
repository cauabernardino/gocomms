package controllers

import (
	"net/http"
	"web/src/utils"
)

// UserLogout handles the loging out of an user by deleting
// the cookie information from the browser
func UserLogout(w http.ResponseWriter, r *http.Request) {

	utils.DeleteCookie(w)
	http.Redirect(w, r, "/login", http.StatusFound)

}
