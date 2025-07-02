/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Rokkit-exe/golly/client"
	"github.com/Rokkit-exe/golly/models"
	"github.com/spf13/cobra"
)

var config = models.LoadConfig("config.yml")

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new custom model",
	Long: `The create command allows you to create a new custom model.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Error getting name flag:", err)
			return
		}
		from, err := cmd.Flags().GetString("from")
		if err != nil {
			fmt.Println("Error getting from flag:", err)
			return
		}
		system, err := cmd.Flags().GetString("system")
		if err != nil {
			fmt.Println("Error getting system flag:", err)
			return
		}

		client := client.NewOllama(config.OllamaHost, config.OllamaPort)

		response, err := client.Create(name, from, system)
		if err != nil {
			fmt.Println("Error creating custom model:", err)
			return
		}

		fmt.Printf("Created successfully \n%s", response.Status)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", config.Model+"-custom", "Name of the custom model")
	createCmd.Flags().StringP("from", "f", config.Model, "Base model to create the custom model from")
	createCmd.Flags().StringP("system", "s", "", "System prompt for the custom model")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("from")
	createCmd.MarkFlagRequired("system")
}
