package main

import (
	"fmt"

	"github.com.br/patricksalmeida/course-go/14-restApi/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
