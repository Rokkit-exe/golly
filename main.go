/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"os"

	"github.com/Rokkit-exe/golly/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("OLLAMA_HOST")
	port := os.Getenv("OLLAMA_PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "11434"
	}

	cmd.Execute()
}
