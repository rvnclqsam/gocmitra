package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
	"github.com/spf13/cobra"
)

var setModelCmd = &cobra.Command{
	Use:   "config set-model [model]",
	Short: "Set the default LLM model",
	Long: `Update the local configuration to use a specific model.
Example: gocmitra config set-model claude-3-sonnet`,
	Args: cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		const cfgPath = ".gocmitra.json"
		model := args[0]

		// Load existing config file
		data, err := os.ReadFile(cfgPath)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to read config file %q: %v", cfgPath, err))
			os.Exit(1)
		}

		var cfg map[string]interface{}
		if err := json.Unmarshal(data, &cfg); err != nil {
			logger.Error(fmt.Sprintf("Config file is not valid JSON: %v", err))
			os.Exit(1)
		}

		// Apply new model
		cfg["model"] = model

		// Save updated config
		newCfg, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to encode config: %v", err))
			os.Exit(1)
		}

		if err := os.WriteFile(cfgPath, newCfg, 0644); err != nil {
			logger.Error(fmt.Sprintf("Failed to save config to %q: %v", cfgPath, err))
			os.Exit(1)
		}

		logger.Success(fmt.Sprintf("Model successfully updated to: %s", model))
	},
}

func init() {
	rootCmd.AddCommand(setModelCmd)
}
