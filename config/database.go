package config

import (
	//"github.com/joho/godotenv"
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
