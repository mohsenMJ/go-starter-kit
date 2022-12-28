package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app"
	"github.com/mohsenMj/go-starter-kit/controllers"
)

func init() {
	app.LoadEnvVariables()
	app.ConnectToDB()
}

var (
	// authController controllers.AuthController = controllers.NewAuthController()

	postController controllers.PostController = controllers.NewPostController()
)

func main() {
	defer app.CloseDbConnection()

	r := gin.Default()

	// r.Use(middlewares.ErrorHandler)

	// authGroup := r.Group("api/auth")
	// {
	// 	authGroup.POST("/login", authController.Login)
	// }

	postGroup := r.Group("/posts")
	{
		postGroup.GET("/", postController.Index)
		postGroup.GET("/:id", postController.Show)
		postGroup.POST("/", postController.Create)
		postGroup.PUT("/:id", postController.Update)
		postGroup.DELETE("/:id", postController.Delete)
	}

	r.Run()
}
