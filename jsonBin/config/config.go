package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key       string
	MasterKey string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error for download .env file: %v", err)
	}

	key := os.Getenv("KEY")
	if key == "" {
		return nil, fmt.Errorf("variable environment KEY don't found")
	}

	masterKey := os.Getenv("X_MASTER_KEY")
	if masterKey == "" {
		return nil, fmt.Errorf("environment variable X-MASTER-KEY not found")
	}

	return &Config{
		Key:       key,
		MasterKey: masterKey,
	}, nil
}
