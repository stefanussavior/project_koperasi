package routes

import (
	"project_koperasi/config"
	"project_koperasi/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routing() {
	routes := gin.Default()

	routes.Use(cors.Default())

	routes.GET("/Coba", config.ConnectDatabase)

	routes.POST("/login", controllers.AuthMidlleware)

	routes.Run(":3000")
}
