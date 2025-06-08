package llm

import (
	"fmt"
	"strings"

	"github.com/Cre4T3Tiv3/gocmitra/core/config"
	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
	"github.com/Cre4T3Tiv3/gocmitra/core/util"
)

// Client defines the interface for any LLM provider.
type Client interface {
	Generate(prompt string, cfg config.Config) (string, error)
	Name() string
}

// NewClient returns the appropriate Client implementation based on the configured endpoint.
// Falls back to OpenAIClient if no specific match is found.
func NewClient(cfg config.Config) Client {
	endpoint := strings.ToLower(cfg.Endpoint)

	switch {
	case strings.Contains(endpoint, "anthropic.com"):
		return ClaudeClient{}
	case strings.Contains(endpoint, "localhost"):
		return OllamaClient{}
	case strings.Contains(endpoint, "openai.com"):
		return OpenAIClient{}
	default:
		logger.Warn(fmt.Sprintf("Unrecognized LLM endpoint: %s – falling back to OpenAI", util.Redact(cfg.Endpoint))) // ✅ patched
		return OpenAIClient{}
	}
}
