package main

import (
	"github.com/PrasGi/learn-golang/controllers"
	"github.com/PrasGi/learn-golang/initializers"
	"github.com/PrasGi/learn-golang/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	route := gin.Default()

	route.GET("/articels", controllers.Index)
	route.POST("/articels", controllers.Store)
	route.DELETE("/articels/:id", controllers.Destroy)
	route.GET("/articels/:id", controllers.Show)
	route.PUT("/articels/:id", controllers.Update)

	route.POST("/signup", controllers.SignUp)
	route.POST("/signin", controllers.Login)

	route.GET("/try", middleware.RequireAuth, controllers.Validate)

	route.Run(":8080")
}
