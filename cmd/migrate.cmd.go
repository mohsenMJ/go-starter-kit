package cmd

import (
	"fmt"

	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/models"
)

func CmdMigrateHandler() {
	fmt.Println("================= Start Migration =================")
	database.DB.Debug().AutoMigrate(&models.Post{}, &models.User{})
	fmt.Println("================= Migration Done =================")
}
