package views

import (
	"net/http"
	"web/src/utils"
)

// LoginPage handles the loading of login page
func LoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := utils.CheckCookie(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	RenderTemplate(w, "login.html", nil)
}
