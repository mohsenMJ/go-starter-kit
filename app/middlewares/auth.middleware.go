package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mohsenMj/go-starter-kit/app/services"
	"github.com/mohsenMj/go-starter-kit/utils"
)

func AuthorizeJWT(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		response := utils.BuildErrorResponse("failed to preocess reqeust", errors.New("No token found"), nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	token, err := services.NewJWTService().ValidateToken(authHeader)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		log.Println(claims)
		log.Println("Claim[user_id]", claims["user_id"])

	} else {
		response := utils.BuildErrorResponse("Token is not validate", err, nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
}
