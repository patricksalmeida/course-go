package main

import (
	"github.com.br/patricksalmeida/course-go/15-restApiGin/models"
	"github.com.br/patricksalmeida/course-go/15-restApiGin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{"Patrick Almeida", "12312312300", "991234567"},
		{"Sabrina Almeida", "12312312400", "991234568"},
	}

	routes.HandleRequests()
}
