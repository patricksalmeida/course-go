package main

import (
	"github.com.br/patricksalmeida/course-go/15-restApiGin/database"
	"github.com.br/patricksalmeida/course-go/15-restApiGin/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
