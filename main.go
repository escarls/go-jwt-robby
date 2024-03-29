package main

import (
	"github.com/escarls/go-jwt-robby/controllers"
	"github.com/escarls/go-jwt-robby/initializers"
	"github.com/escarls/go-jwt-robby/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
