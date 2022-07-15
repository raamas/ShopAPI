package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err = godotenv.Load()
var DB *gorm.DB

var password = os.Getenv("DB_PASSWORD")
var host = os.Getenv("DB_HOST")
var port = os.Getenv("DB_PORT")
var dbString = fmt.Sprintf("user=postgres password=%v host=%v port=%v dbname=postgres", password, host, port)

func Connect() {
	var error error
	DB, error = gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connection successful")
	}
}
