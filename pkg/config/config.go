package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("error load env : %v", err)
	}
	return nil
}
