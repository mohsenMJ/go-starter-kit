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

	// public routes
	postGroup := r.Group("/posts")
	{
		postGroup.GET("/", postController.PostIndex)
		postGroup.GET("/:id", postController.PostShow)
		postGroup.POST("/", postController.PostCreate)
		postGroup.PUT("/:id", postController.PostUpdate)
		postGroup.DELETE("/:id", postController.PostDelete)
	}
}
