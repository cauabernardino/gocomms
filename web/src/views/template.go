package views

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadTemplates inserts the html templates into the variable
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

//
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
