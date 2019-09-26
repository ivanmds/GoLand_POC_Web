package controllers

import (
	"GoLand_POC_Web/models"
	"html/template"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "None name sended"
	}

	person := models.Person{Name: name}
	renderTemplate(w, "./GoLand_POC_Web/templates/greeting.html", person)
}

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}
