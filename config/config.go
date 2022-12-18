package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUrl string
}

func LoadConfig() (config *Config) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	mongoUrl := os.Getenv("MONGODB_URL")
	return &Config{
		MongoUrl: mongoUrl,
	}
}
