/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/Rokkit-exe/golly/config"
	"github.com/Rokkit-exe/golly/models"
	"github.com/Rokkit-exe/golly/ollama"
	"github.com/Rokkit-exe/golly/ui"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var query string

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with Ollama",
	Long: `The chat command allows you to interact with Ollama's chat capabilities.à
You can specify the model, host, and port to connect to your Ollama instance.
- Example usage:
	golly chat --model llama3.2 --host localhost --port 11434`,
	Run: func(cmd *cobra.Command, args []string) {
		model, err := cmd.Flags().GetString("model")
		if err != nil {
			fmt.Println("Error getting model flag:", err)
		}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			fmt.Println("Error getting host flag:", err)
		}

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			fmt.Println("Error getting port flag:", err)
		}
		query := strings.Join(args, " ")
		fmt.Println("model: " + model)
		fmt.Println("host: " + host)
		fmt.Println("port: " + port)
		fmt.Println("query: " + query)

		time.Sleep(2 * time.Second) // Optional: small delay for better UX

		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
		)
		if err != nil {
			fmt.Println("Error creating renderer:", err)
			return
		}
		ui := ui.UI{
			Renderer:     renderer,
			Query:        "",
			FullResponse: "",
		}
		ollamaClient := ollama.NewOllama(host, port)
		for quit := false; !quit; {
			streamCh, errCh := ollamaClient.StreamChat(model, []models.ChatMessage{
				{Role: "user", Content: query},
			})
			ui.PrintAI(streamCh, errCh)
			query, ok := ui.Scan()
			if !ok {
				quit = true
				continue
			}
			ui.PrintUser(query)
		}
	},
}

func init() {
	config, err := config.LoadConfig(".env")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().StringVar(&query, "query", "Hello!", "Query to send to the chat model")
	chatCmd.Flags().StringP("model", "m", config.DefaultModel, "Model to use for the chat")
	chatCmd.Flags().StringP("host", "H", config.Host, "Host of the Ollama instance")
	chatCmd.Flags().StringP("port", "p", config.Port, "Port of the Ollama instance")
}
