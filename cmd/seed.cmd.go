package cmd

import (
	"fmt"

	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/seeders"
)

func init() {
	Commands = append(Commands, Command{
		Run:         "seed",
		Description: "seed database",
		Example:     "go run . seed",
		Author:      "Mohsen Majidi",
	})
}

func CmdRunSeedersHandler() {
	fmt.Println("================= Run Seeders =================")
	seeders.UsersSeeder(database.DB)
	// add your seeders here
	fmt.Println("================= Seeders Added =================")
}
