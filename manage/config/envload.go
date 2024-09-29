package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoad()string{
	err:= godotenv.Load()
	if err !=nil{
		log.Fatal("Error loading .env files")
	}

	return os.Getenv("Mongo")
}