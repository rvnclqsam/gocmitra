package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Cre4T3Tiv3/gocmitra/core/config"
	"github.com/Cre4T3Tiv3/gocmitra/core/diff"
	"github.com/Cre4T3Tiv3/gocmitra/core/llm/providers"
	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
	"github.com/Cre4T3Tiv3/gocmitra/core/prompt"
	"github.com/Cre4T3Tiv3/gocmitra/core/util" // <-- added
	"github.com/spf13/cobra"
)

var staged bool

func init() {
	diffCmd.Flags().BoolVar(&staged, "staged", false, "Use staged changes (git diff --staged)")
	rootCmd.AddCommand(diffCmd)
}

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Generate a commit message from Git diff",
	Run: func(_ *cobra.Command, _ []string) {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigs
			fmt.Println("\n[⚠️] Aborted.")
			os.Exit(130)
		}()

		gitArgs := []string{"diff"}
		if staged {
			gitArgs = append(gitArgs, "--staged")
		}
		out, err := exec.Command("git", gitArgs...).Output()
		if err != nil {
			logger.Error(fmt.Sprintf("git diff failed: %v", err))
			os.Exit(1)
		}
		logger.Success("Git diff collected")

		cfg, err := config.Load(cfgPath)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to load config (%s): %v", cfgPath, err))
			os.Exit(1)
		}
		logger.Success("Loaded config")

		if envModel := os.Getenv("GOCMITRA_MODEL"); envModel != "" {
			cfg.Model = envModel
		}
		if model != "" {
			cfg.Model = model
		}

		logger.Info(fmt.Sprintf("Current model: %s @ %s", cfg.Model, util.Redact(cfg.Endpoint)))

		var userResponse string
		for {
			logger.Info("Would you like to change the model? (y/N): ")
			fmt.Print("> ")
			_, err := fmt.Scanln(&userResponse)
			if err != nil {
				continue
			}
			userResponse = strings.ToLower(strings.TrimSpace(userResponse))
			if userResponse == "y" || userResponse == "n" || userResponse == "" {
				break
			}
			logger.Warn("Invalid input. Please enter 'y' or 'n'.")
		}

		if userResponse == "y" {
			logger.Info("Select a model:")
			logger.Info("1. gpt-4o (OpenAI)")
			logger.Info("2. claude-3 (Anthropic)")
			logger.Info("3. llama3 (Ollama)")

			var choice int
			for {
				fmt.Print("> ")
				_, err := fmt.Scanln(&choice)
				if err != nil {
					logger.Warn("Invalid input. Please enter a number (1–3).")
					continue
				}

				var profilePath string
				switch choice {
				case 1:
					profilePath = "profiles/.gocmitra-openai.json"
				case 2:
					profilePath = "profiles/.gocmitra-claude3.json"
				case 3:
					profilePath = "profiles/.gocmitra-llama3.json"
				default:
					logger.Warn("Invalid selection. Please enter 1, 2, or 3.")
					continue
				}

				profileCfg, err := config.Load(profilePath)
				if err != nil {
					logger.Warn(fmt.Sprintf("Failed to load model profile: %v", err))
				} else {
					cfg = profileCfg
					logger.Info(fmt.Sprintf("Model config updated: %s @ %s", cfg.Model, util.Redact(cfg.Endpoint)))

					if err := persistConfig(cfgPath, cfg); err != nil {
						logger.Warn(fmt.Sprintf("Failed to persist config: %v", err))
					}
				}
				break
			}
		}

		logger.Info("Parsing git diff content")
		diffs := diff.Parse(string(out))
		promptContent := prompt.Build(diffs, cfg)

		client := llm.NewClient(cfg)
		logger.Success(fmt.Sprintf("Using model: %s @ %s", cfg.Model, client.Name()))
		logger.Success("Sending prompt to model...")

		result, err := client.Generate(promptContent, cfg)
		if err != nil {
			logger.Error(fmt.Sprintf("LLM error: %v", err))
			os.Exit(1)
		}
		logger.Success("Response received:")
		fmt.Println()

		clean := strings.TrimSpace(result)
		if strings.HasPrefix(clean, "```") {
			if parts := strings.SplitN(clean, "\n", 2); len(parts) == 2 {
				clean = strings.TrimSuffix(parts[1], "```")
				clean = strings.TrimSpace(clean)
			}
		}

		fmt.Println(clean)
	},
}

func persistConfig(path string, cfg config.Config) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(cfg)
}
