package main

import (
	"github.com.br/patricksalmeida/course-go/16-unitaryTests/database"
	"github.com.br/patricksalmeida/course-go/16-unitaryTests/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
