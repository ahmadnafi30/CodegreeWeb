package config

import (
	"log"
	"os"

	// "os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	env := os.Getenv("ENV")
	if err != nil && env == "" {
		log.Fatalf("error load env : %v", err)
	}
	return nil
}
