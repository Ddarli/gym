package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var HttpAddr string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	HttpAddr = os.Getenv("HTTP_ADDRESS")
}
