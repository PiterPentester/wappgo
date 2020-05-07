package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates - parse & load templates to server
func LoadTemplates(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))
}

// ExecuteTemplate - execute your template on webserver
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}
