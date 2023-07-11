package routes

import (
	"github.com.br/patricksalmeida/course-go/15-restApiGin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/:nome", controllers.Saudacao)

	r.Run()
}
