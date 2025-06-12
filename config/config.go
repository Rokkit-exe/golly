package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	DefaultModel string `json:"default_model"`
}

func LoadConfig(filePath string) (*Config, error) {
	// Load environment variables from the specified filePath
	err := godotenv.Load(filePath)
	if err != nil {
		return nil, err
	}

	host := os.Getenv("OLLAMA_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("OLLAMA_PORT")
	if port == "" {
		port = "11434"
	}

	model := os.Getenv("DEFAULT_MODEL")
	if model == "" {
		model = "llama3.2"
	}

	config := &Config{
		Host:         host,
		Port:         port,
		DefaultModel: model,
	}

	// Here you would typically load the config from a file or environment variables.
	// For simplicity, we are returning the default config.
	return config, nil
}
