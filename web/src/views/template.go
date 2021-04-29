package views

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadTemplates inserts the html templates into the variable
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	templates = template.Must(templates.ParseGlob("templates/base/*.html"))
}

// RenderTemplate renders the templates to the browser
func RenderTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
