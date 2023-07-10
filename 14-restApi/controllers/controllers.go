package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com.br/patricksalmeida/course-go/14-restApi/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
