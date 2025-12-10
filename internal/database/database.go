package database

import (
	"fmt"
	"go_project_Gin/internal/config"
	"go_project_Gin/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Dhaka",
		config.App.DBHost,
		config.App.DBUser,
		config.App.DBPassword,
		config.App.DBName,
		config.App.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB = db
	log.Println("Database connected")
	if err := DB.AutoMigrate(&model.Task{}, &model.User{}); err != nil {
		log.Fatal("Failed to migrate database", err)
	}
	log.Println("Database migrated")
}
