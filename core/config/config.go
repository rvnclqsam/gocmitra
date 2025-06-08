package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
)

// Config holds application configuration loaded from .gocmitra.json.
type Config struct {
	Model          string `json:"model"`
	Style          string `json:"style"`
	Tone           string `json:"tone"`
	Instructions   string `json:"instructions"`
	TicketPattern  string `json:"ticketPattern"`
	Endpoint       string `json:"endpoint"`
	PromptTemplate string `json:"promptTemplate,omitempty"`
	APIKey         string `json:"apiKey,omitempty"` // Injected during config load
}

// Default returns a Config with sensible defaults.
func Default() Config {
	return Config{
		Model:         "gpt-4o",
		Style:         "conventional",
		TicketPattern: "",
		Endpoint:      "https://api.openai.com/v1/chat/completions",
	}
}

// Load reads configuration from path. If the file does not exist,
// Default configuration is returned.
func Load(path string) (Config, error) {
	cfg := Default()

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return injectAPIKey(cfg), nil
		}
		logger.Error(fmt.Sprintf("Failed to open config file %s: %v", path, err))
		return cfg, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		logger.Error(fmt.Sprintf("Failed to decode config file %s: %v", path, err))
		return cfg, err
	}

	return injectAPIKey(cfg), nil
}

// injectAPIKey resolves the appropriate environment variable based on the model and injects it into config.
func injectAPIKey(cfg Config) Config {
	model := strings.ToLower(cfg.Model)

	switch {
	case strings.Contains(model, "gpt"):
		cfg.APIKey = os.Getenv("OPENAI_API_KEY")
	case strings.Contains(model, "claude"):
		cfg.APIKey = os.Getenv("CLAUDE_API_KEY")
	case strings.Contains(model, "llama"):
		// Ollama is local, no API key needed
	default:
		cfg.APIKey = os.Getenv("GOCMITRA_API_KEY")
	}

	return cfg
}
