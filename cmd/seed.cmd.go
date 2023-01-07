package cmd

import (
	"fmt"

	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/seeders"
)

func CmdRunSeedersHandler() {
	fmt.Println("================= Run Seeders =================")
	seeders.UsersSeeder(database.DB)
	// add your seeders here
	fmt.Println("================= Seeders Added =================")
}
