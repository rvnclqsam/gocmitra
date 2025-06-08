package cmd

import (
	"fmt"

	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
	"github.com/spf13/cobra"
)

var (
	cfgPath string
	model   string

	// Define the CLI version here
	Version = "v0.1.0-beta"
)

// rootCmd defines the base command for gocmitra
var rootCmd = &cobra.Command{
	Use:     "gocmitra",
	Short:   "GoCmitra is an AI-powered Git commit assistant",
	Long:    `GoCmitra analyzes Git diffs and uses LLMs to suggest intelligent commit messages.`,
	Version: Version,
}

// Execute runs the root command and handles any top-level execution errors.
func Execute() error {
	// Initialize logger
	logger.Init()

	// Register completion command
	rootCmd.AddCommand(completionCmd)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", ".gocmitra.json", "Path to config file")
	rootCmd.PersistentFlags().StringVar(&model, "model", "", "Override model from config")

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		logger.Error(fmt.Sprintf("CLI execution failed: %v", err))
		return err
	}
	return nil
}
