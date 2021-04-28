package views

import "net/http"

// SignUpPage handles the loading of sign up page
func SignUpPage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "signup.html", nil)
}
