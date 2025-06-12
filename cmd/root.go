/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Rokkit-exe/golly/models"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var (
	model string
	host  string
	port  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golly",
	Short: "Golly CLI",
	Long:  `Golly is a command-line interface for managing and interacting with Ollama with style in the terminal.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		model, err := cmd.Flags().GetString("model")
		if err != nil || model == "" {
			model = "llama3.2" // Default model
		}
		host, err := cmd.Flags().GetString("host")
		if err != nil || host == "" {
			host = "localhost" // Default host
		}

		port, err := cmd.Flags().GetString("port")
		if err != nil || port == "" {
			port = "11434" // Default port
		}
		fmt.Println("model: " + model)
		fmt.Println("host: " + host)
		fmt.Println("port: " + port)

		reqBody := models.ChatRequest{
			Model: model,
			Messages: []models.ChatMessage{
				{Role: "user", Content: "how can i write a function in python that prints 'hello world'?"},
			},
			Stream: true,
		}

		jsonData, err := json.Marshal(reqBody)
		if err != nil {
			fmt.Println("Error marshaling request body:", err)
			return
		}

		resp, err := http.Post(
			"http://"+host+":"+port+"/api/chat",
			"application/json",
			bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error making POST request:", err)
			return
		}
		defer resp.Body.Close()

		renderer, _ := glamour.NewTermRenderer(glamour.WithAutoStyle())
		scanner := bufio.NewScanner(resp.Body)

		var buf bytes.Buffer
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				fmt.Println("Received empty line, skipping...")
				continue
			}
			var chunk models.ChatResponseChunk
			if err := json.Unmarshal([]byte(line), &chunk); err != nil {
				fmt.Printf("JSON unmarshal error: %v (%q)\n", err, line)
				continue
			}

			buf.WriteString(chunk.Message.Content)

			rendered, err := renderer.Render(buf.String())
			if err != nil {
				fmt.Printf("Glamour render error: %v\n", err)
				continue
			}

			fmt.Print("\033[H\033[2J") // Optional: clears screen for live update
			fmt.Print(rendered)

			if chunk.Done {
				break
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Scanner err: %v", err)
		}

		fmt.Println("\nStreaming complete.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golly.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("model", "m", "llama3.2", "Model to use for Ollama (default: llama3.2)")
	rootCmd.Flags().StringP("port", "p", "11434", "Port to use for Ollama server (default: 11434)")
	rootCmd.Flags().StringP("host", "H", "localhost", "Host to use for Ollama server (default: localhost)")
}
