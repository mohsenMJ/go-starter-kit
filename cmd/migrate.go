package main

import (
	"fmt"

	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/initializers/env"
	"github.com/mohsenMj/go-starter-kit/app/models"
)

func init() {
	env.LoadVariables()
	database.Connect()
}

func main() {
	defer database.Disconnect()

	fmt.Println("Starting to migrate the database")

	database.DB.Debug().AutoMigrate(&models.Post{}, &models.User{})

	fmt.Println("Migration is Done")

}
