package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rishi058/go-project/models"
	"github.com/rishi058/go-project/storage"
	"gorm.io/gorm"
)


var Instance *gorm.DB

func InitializeDB() {

	err := godotenv.Load(".env")

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	
	db, err := storage.NewConnection(config)  // establishing connection with the postgres DB server.

	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateData(db)   // creating tables if not exist.

	if err != nil {
		log.Fatal("could not migrate db")
	}

	Instance = db;
}