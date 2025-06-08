package llm

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Cre4T3Tiv3/gocmitra/core/config"
	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
)

type OllamaClient struct{}

func (o OllamaClient) Generate(prompt string, cfg config.Config) (string, error) {
	systemMsg := cfg.Instructions

	body := map[string]interface{}{
		"model": cfg.Model,
		"messages": []map[string]string{
			{"role": "system", "content": systemMsg},
			{"role": "user", "content": prompt},
		},
	}

	raw, err := sendChatRequest(cfg.Endpoint, body, nil)
	if err != nil {
		return "", fmt.Errorf("Ollama request failed: %w", err)
	}

	clean := strings.TrimSpace(raw)
	if clean == "" {
		return "", fmt.Errorf("Ollama returned an empty response")
	}

	// Try to parse JSON if response is structured
	var parsed struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	}
	if err := json.Unmarshal([]byte(clean), &parsed); err == nil && parsed.Message.Content != "" {
		clean = parsed.Message.Content
	} else if strings.HasPrefix(clean, "{") {
		logger.Warn(fmt.Sprintf("Ollama response might be malformed JSON: %s", clean))
	}

	// Normalize output
	line := strings.SplitN(clean, "\n", 2)[0]
	line = strings.TrimSpace(line)
	line = strings.Trim(line, "`\".")
	line = strings.ReplaceAll(line, "\"", "")

	if i := strings.Index(line, ":"); i != -1 && len(line) > i+1 {
		prefix := line[:i+1]
		rest := strings.ToLower(strings.TrimSpace(line[i+1:]))
		line = prefix + " " + rest
	}

	if line == "" {
		return "", fmt.Errorf("Ollama response parsed but resulted in empty commit message")
	}

	return line, nil
}

func (o OllamaClient) Name() string {
	return "Ollama"
}
