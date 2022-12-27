package main

import (
	"go-curd/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Load()
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
