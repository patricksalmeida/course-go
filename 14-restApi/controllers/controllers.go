package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com.br/patricksalmeida/course-go/14-restApi/database"
	"github.com.br/patricksalmeida/course-go/14-restApi/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade

	database.DB.Find(&personalidades)

	json.NewEncoder(w).Encode(personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}
