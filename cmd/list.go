/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Rokkit-exe/golly/client"
	"github.com/Rokkit-exe/golly/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available models",
	Long: `List all available models from the Ollama server.
Custom and default models will be displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := client.NewOllama(config.Host, config.Port)

		resp, err := client.List()
		if err != nil {
			fmt.Println("Error listing models:", err)
			return
		}

		fmt.Println("Available models:")
		if len(resp.Models) == 0 {
			fmt.Println("No models found.")
			return
		}
		for _, model := range resp.Models {
			utils.PrintStruct(model)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
