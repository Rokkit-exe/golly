package models

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OllamaHost    string        `yaml:"ollama_host"`
	OllamaPort    string        `yaml:"ollama_port"`
	Model         string        `yaml:"model"`
	SearxngHost   string        `yaml:"searxng_host"`
	SearxngPort   string        `yaml:"searxng_port"`
	SystemPrompts SystemPrompts `yaml:"system_prompts"`
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
		OllamaHost:  host,
		OllamaPort:  port,
		Model:       model,
		SearxngHost: host,
		SearxngPort: "8080",
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

	if config.OllamaHost == "" {
		config.OllamaHost = host
	}
	if config.OllamaPort == "" {
		config.OllamaPort = port
	}
	if config.Model == "" {
		config.Model = model
	}
	if config.SearxngHost == "" {
		config.SearxngHost = host
	}
	if config.SearxngPort == "" {
		config.SearxngPort = "8080"
	}
	fmt.Println("Configuration loaded successfully: host:",
		config.OllamaHost, "port:",
		config.OllamaPort, "model:",
		config.Model, "searxng_host:",
		config.SearxngHost, "searxng_port:",
		config.SearxngPort,
	)

	return config
}
