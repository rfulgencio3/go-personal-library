package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string
	MongoURI        string
	MongoDatabase   string
	MongoCollection string
}

func LoadConfig() (*Config, error) {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := &Config{
		ServerPort:      os.Getenv("SERVER_PORT"),
		MongoURI:        os.Getenv("MONGO_URI"),
		MongoDatabase:   os.Getenv("MONGO_DATABASE"),
		MongoCollection: os.Getenv("MONGO_COLLECTION"),
	}

	return config, nil
}
