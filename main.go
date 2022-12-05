package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maurodesouza/jwt-auth-golang/controllers"
	"github.com/maurodesouza/jwt-auth-golang/initializers"
	"github.com/maurodesouza/jwt-auth-golang/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.GET("/validate", middlewares.AuthMiddleware, controllers.Validate)

	r.Run()
}
