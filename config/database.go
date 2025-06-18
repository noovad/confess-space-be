package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func DatabaseConnection() *gorm.DB {
	var err error

	var dsn string
	if os.Getenv("ENV") == "production" {
		fmt.Println("Database connection in production mode")
		dsn = os.Getenv("DATABASE_URL")
	} else {
		host := os.Getenv("DBHOST")
		user := os.Getenv("DBUSER")
		password := os.Getenv("DBPASSWORD")
		dbname := os.Getenv("DBNAME")
		port := os.Getenv("DBPORT")

		if host == "" || user == "" || password == "" || dbname == "" || port == "" {
			log.Fatal("One or more required environment variables are not set")
		}
		fmt.Println("Database connection in development mode")
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			host, user, password, dbname, port)
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	return db
}
