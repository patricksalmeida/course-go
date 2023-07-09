package main

import (
	"net/http"

	"github.com.br/patricksalmeida/course-go/13-webapp/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
