package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app"
	"github.com/mohsenMj/go-starter-kit/controllers"
	"github.com/mohsenMj/go-starter-kit/middlewares"
)

func init() {
	app.LoadEnvVariables()
	app.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.Use(middlewares.ErrorHandler)

	r.POST("/posts", controllers.PostCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
