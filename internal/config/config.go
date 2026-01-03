package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	DBHost              string
	DBPort              int
	DBUser              string
	DBPassword          string
	DBName              string
	JWTSecret           string
	AppName             string
	Environment         string
	FROM_EMAIL          string
	FROM_EMAIL_PASSWORD string
}

var App Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	App = Config{
		Port:                getEnv("PORT", "8080"),
		DBHost:              getEnv("DB_HOST", "localhost"),
		DBUser:              getEnv("DB_USER", "go_user"),
		DBPassword:          getEnv("DB_PASSWORD", "123456"),
		DBName:              getEnv("DB_NAME", "taskdb"),
		JWTSecret:           getEnv("JWT_SECRET", "fallback-secret-key"),
		AppName:             getEnv("APP_NAME", "TaskAPI"),
		Environment:         getEnv("ENV", "development"),
		FROM_EMAIL:          getEnv("FROM_EMAIL", "islamrafiq9090@gmail.com"),
		FROM_EMAIL_PASSWORD: getEnv("FROM_EMAIL_PASSWORD", "dwupsrscfczakuxc"),
	}
	if port, err := strconv.Atoi(getEnv("DB_PORT", "5432")); err == nil {
		App.DBPort = port
	} else {
		App.DBPort = 5432
	}
	log.Printf("Config loaded – Running in %s mode\n", App.Environment)
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
