package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectToDatabase() {
	dsn := "myuser:mypassword@tcp(localhost:3306)/mydatabase?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	database = db
}

func MigrateDatabase() {
	database.AutoMigrate(Note{})
}
