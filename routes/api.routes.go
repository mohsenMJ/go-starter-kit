package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app/http/controllers"
)

var (
	authController controllers.AuthController = controllers.NewAuthController()
	postController controllers.PostController = controllers.NewPostController()
)

func Api(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Go Starter Kit",
		})
	})

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
