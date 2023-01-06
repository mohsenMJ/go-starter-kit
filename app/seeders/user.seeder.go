package seeders

import (
	"fmt"

	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/models"
	"gorm.io/gorm"
)

func UsersSeeder(db *gorm.DB) {
	fmt.Print("User Seeder: ")
	users := []models.User{
		models.User{Name: "Mohsen", Password: "12345678", Email: "mohsen.majidi22@gmail.com"},
		models.User{Name: "Mohsen2", Password: "12345678", Email: "mohsen.majidi23@gmail.com"},
	}
	for _, user := range users {
		database.DB.Save(&user)
	}
	fmt.Println("Done")
}
