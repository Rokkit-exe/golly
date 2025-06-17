package models

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host          string   `yaml:"host"`
	Port          string   `yaml:"port"`
	Model         string   `yaml:"model"`
	SystemPrompts []string `yaml:"system_prompts"`
}

type SystemPrompts struct {
	RefactorQuery string `yaml:"refactor_query"`
	ResultPicker  string `yaml:"result_picker"`
	AnswerBuilder string `yaml:"answer_builder"`
}

var (
	configPath = "config.yml"
	host       = "localhost"
	port       = "11434"
	model      = "llama3.2"
)

func LoadConfig(filePath string) *Config {
	if filePath != "" {
		configPath = filePath
	}

	defaultConfig := &Config{
		Host:  host,
		Port:  port,
		Model: model,
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		fmt.Println("Using default configuration: host:", host, "port:", port, "model:", model)
		return defaultConfig
	}

	config := &Config{}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		fmt.Println("Using default configuration: host:", host, "port:", port, "model:", model)
		return defaultConfig
	}

	if config.Host == "" {
		config.Host = host
	}
	if config.Port == "" {
		config.Port = port
	}
	if config.Model == "" {
		config.Model = model
	}
	fmt.Println("Configuration loaded successfully: host:", config.Host, "port:", config.Port, "model:", config.Model)

	return config
}
