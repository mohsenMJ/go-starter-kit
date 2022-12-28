package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	log.Println("===================================== middleware - error handler ===============================")
	for _, err := range c.Errors {
		// log, handle, etc.
		log.Println(err)
	}
	log.Println("===================================== middleware - error handler ===============================")

	c.JSON(http.StatusInternalServerError, "")
}
