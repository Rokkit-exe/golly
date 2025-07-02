/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Rokkit-exe/golly/client"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a model",
	Long: `Delete a model from the Ollama server.
You can specify the model name to delete it from the server.
- Example usage:
golly delete --model llama3.2`,
	Run: func(cmd *cobra.Command, args []string) {
		model, err := cmd.Flags().GetString("model")
		if err != nil {
			fmt.Println("Error getting model flag:", err)
			return
		}

		if model == "" {
			fmt.Println("Model name is required. Use --model to specify the model to delete.")
			return
		}

		client := client.NewOllama(config.OllamaHost, config.OllamaPort)

		err = client.Delete(model)
		if err != nil {
			fmt.Println("Error deleting model:", err)
			return
		}

		fmt.Printf("Model '%s' deleted successfully.\n", model)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("model", "m", "", "Name of the model to delete")
	deleteCmd.MarkFlagRequired("model")
}
