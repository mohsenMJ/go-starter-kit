package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// log.Println("Databse is connected Successfuly...")
}

func Disconnect() {
	dbInstance, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to close the connection")
	}
	_ = dbInstance.Close()

	// log.Println("Databse connection closed Successfuly...")
}
