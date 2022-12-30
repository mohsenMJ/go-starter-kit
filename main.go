package main

import (
	"github.com/gin-gonic/gin"
	app "github.com/mohsenMj/go-starter-kit/app/providers"
	"github.com/mohsenMj/go-starter-kit/routes"
)

func init() {
	app.LoadEnvVariables()
	app.DatabaseConnect()
}

func main() {
	defer app.DatabaseDisconnect()
	r := gin.Default()
	routes.Api(r)
	r.Run()
}
