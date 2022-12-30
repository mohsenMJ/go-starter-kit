package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app/services"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	//service goes here
}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	jwt := services.NewJWTService()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Login",
		"token":   jwt.GenerateToken("1"),
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Register",
	})
}
