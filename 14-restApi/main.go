package main

import (
	"fmt"

	"github.com.br/patricksalmeida/course-go/14-restApi/models"
	"github.com.br/patricksalmeida/course-go/14-restApi/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "História 1"},
		{Nome: "Nome 2", Historia: "História 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
