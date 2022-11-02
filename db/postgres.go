package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	error_env := godotenv.Load()
	if error_env != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	path := os.Getenv("DB_PATH")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s search_path=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, path, pass, name, port)

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, error
}
