package main

import (
	"fmt"

	"github.com.br/patricksalmeida/course-go/14-restApi/database"
	"github.com.br/patricksalmeida/course-go/14-restApi/models"
	"github.com.br/patricksalmeida/course-go/14-restApi/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Nome 1", Historia: "História 1"},
		{ID: 2, Nome: "Nome 2", Historia: "História 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
