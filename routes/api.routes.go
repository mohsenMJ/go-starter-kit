package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app/controllers"
)

func Api(r *gin.Engine) {
	authController := controllers.NewAuthController()
	postController := controllers.NewPostController()

	apiGroup := r.Group("/api")
	{
		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/login", authController.Login)
			authGroup.POST("/register", authController.Register)
		}
	}

	postGroup := r.Group("/posts")
	{
		postGroup.GET("/", postController.Index)
		postGroup.GET("/:id", postController.Show)
		postGroup.POST("/", postController.Create)
		postGroup.PUT("/:id", postController.Update)
		postGroup.DELETE("/:id", postController.Delete)
	}
}
