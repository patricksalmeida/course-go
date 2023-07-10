package routes

import (
	"log"
	"net/http"

	"github.com.br/patricksalmeida/course-go/14-restApi/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
