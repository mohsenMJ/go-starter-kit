package initializers

import "github.com/joho/godotenv"

func Load() {
	err := godotenv.Load()
	if err != nil {
		log
		.Fatal("Error loading .env file")
	}
}
