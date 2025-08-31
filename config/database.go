package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {

	// err = godotenv.Load("config/.env")
	// if err != nil {
	// 	panic("Error loading .env file")
	// }

	// psqlinfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )
	psqlinfo := os.Getenv("DATABASE_PUBLIC_URL")
	if psqlinfo == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	database, err := gorm.Open(postgres.Open(psqlinfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to : ", err)
	}

	DB = database

	fmt.Println("Succesfully Connected to Database")
}
