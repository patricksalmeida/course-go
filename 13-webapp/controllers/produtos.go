package controllers

import (
	"html/template"
	"net/http"

	"github.com.br/patricksalmeida/course-go/13-webapp/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()

	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}
