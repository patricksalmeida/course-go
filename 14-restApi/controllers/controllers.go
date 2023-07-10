package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com.br/patricksalmeida/course-go/14-restApi/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.ID) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
