package views

import "net/http"

func LoadSignUpPage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "signup.html", nil)
}
