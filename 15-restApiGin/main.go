package main

import (
	"github.com.br/patricksalmeida/course-go/15-restApiGin/database"
	"github.com.br/patricksalmeida/course-go/15-restApiGin/models"
	"github.com.br/patricksalmeida/course-go/15-restApiGin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Patrick Almeida", CPF: "12312312300", RG: "991234567"},
		{Nome: "Sabrina Almeida", CPF: "12312312400", RG: "991234568"},
	}

	routes.HandleRequests()
}
