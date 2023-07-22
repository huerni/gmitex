package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gmctl",
	Short: "Build a CLI (Command Line Interface) tool for microservice architecture.",
	Long:  `Build a CLI (Command Line Interface) tool for microservice architecture.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
