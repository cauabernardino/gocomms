package views

import (
	"net/http"
)

// LoginPage handles the loading of login page
func LoginPage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login.html", nil)
}
