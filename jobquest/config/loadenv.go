package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load ENV Variables")
	}
	return os.Getenv("MONGODB_URL")
}
