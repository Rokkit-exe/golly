/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start ollama server",
	Long:  `Start the ollama server`,
	Run: func(cmd *cobra.Command, args []string) {
		c := []string{"ollama", "serve", "&"}
		command := exec.Command(c[0], c[1:]...)
		err := command.Start()
		if err != nil {
			cmd.Println("Error starting ollama server:", err)
			return
		}
		cmd.Println("Ollama server started successfully.")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
