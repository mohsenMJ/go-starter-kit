package main

import (
	"fmt"
	"go-curd/app"
	"go-curd/models"
)

func init() {
	app.LoadEnvVariables()
	app.ConnectToDB()
}

func main() {
	fmt.Println("Starting to migrate the database")

	app.DB.Debug().AutoMigrate(&models.Post{})

	fmt.Println("Migration is Done")

}
