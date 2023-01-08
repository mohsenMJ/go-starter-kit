package cmd

import (
	"fmt"

	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/models"
)

func init() {
	Commands = append(Commands, Command{
		Run:         "migrate",
		Description: "migrate database",
		Example:     "go run . migrate",
		Author:      "Mohsen Majidi",
	})
}
func CmdMigrateHandler() {
	fmt.Println("================= Start Migration =================")
	database.DB.Debug().AutoMigrate(&models.Post{}, &models.User{})
	fmt.Println("================= Migration Done =================")
}
