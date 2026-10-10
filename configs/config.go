package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	EnvFile string
}

func (e *EnvConfig) GetDbUrl() string {
	err := godotenv.Load(e.EnvFile)
	if err != nil {
		log.Fatal("Error loading .env file...")
	}
	DbUrl := os.Getenv("DB_URL")
	return DbUrl
}
