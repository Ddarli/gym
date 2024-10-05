package common

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     int
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &Config{
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     port,
	}
}
