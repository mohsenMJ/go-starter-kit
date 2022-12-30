package main

import (
	"fmt"

	app "github.com/mohsenMj/go-starter-kit/app/initializers"
	"github.com/mohsenMj/go-starter-kit/database/models"
)

func init() {
	app.LoadEnvVariables()
	app.DatabaseConnect()
}

func main() {
	defer app.DatabaseDisconnect()

	fmt.Println("Starting to migrate the database")

	app.DB.Debug().AutoMigrate(&models.Post{}, &models.User{})

	fmt.Println("Migration is Done")

}
