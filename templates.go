package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.New("").Funcs(templateFunctions).ParseFiles("templates/header.html",
	"./templates/dashboard.html",
	"./templates/footer.html"))

// Register the functions required to render the templates.
var templateFunctions = template.FuncMap{}

// Helper function to render templates.
func renderTemplate(w http.ResponseWriter, templateName string, model interface{}) {
	err := templates.ExecuteTemplate(w, templateName, model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
