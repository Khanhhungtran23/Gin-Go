package config

import (
	"fmt"
	"log"
	"myproject/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// connect database
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "1234")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "gin-database")

	dsn := os.Getenv("DB_USERNAME") + ":" +
		os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println("Database DSN:", dsn) // use this dsn to connect database

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&models.Book{})
}
