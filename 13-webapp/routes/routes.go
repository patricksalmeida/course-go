package routes

import (
	"net/http"

	"github.com.br/patricksalmeida/course-go/13-webapp/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
