/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/initializers/env"
	"github.com/mohsenMj/go-starter-kit/cmd"
	"github.com/mohsenMj/go-starter-kit/routes"
)

func init() {
	env.LoadVariables()
	database.Connect()
}

func main() {
	defer database.Disconnect()
	if len(os.Args) > 1 {
		cmd.Execute(os.Args)
	} else {
		r := gin.Default()
		routes.Api(r)
		r.Run()
	}

}
