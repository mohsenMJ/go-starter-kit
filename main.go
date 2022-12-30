package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/initializers/env"
	"github.com/mohsenMj/go-starter-kit/routes"
)

func init() {
	env.LoadVariables()
	database.Connect()

	log.Println("-----------------------------", database.DB)
}

func main() {
	defer database.Disconnect()
	r := gin.Default()
	routes.Api(r)
	r.Run()
}
