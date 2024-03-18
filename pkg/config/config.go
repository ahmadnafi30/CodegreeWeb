package config

import (
	"log"
	// "os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("error load env : %v", err)
	}
}
