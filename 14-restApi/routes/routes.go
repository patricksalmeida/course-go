package routes

import (
	"log"
	"net/http"

	"github.com.br/patricksalmeida/course-go/14-restApi/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades)

	log.Fatal(http.ListenAndServe(":8080", r))
}
