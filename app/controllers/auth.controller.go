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
	jwt services.JWTService
}

func NewAuthController() AuthController {
	return &authController{
		jwt: services.NewJWTService(),
	}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Login",
		"token":   c.jwt.GenerateToken("1"),
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":       "Hello Register",
		"authenticated": "Test",
	})
}
