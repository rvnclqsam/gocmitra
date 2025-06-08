package llm

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Cre4T3Tiv3/gocmitra/core/config"
	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
	"github.com/Cre4T3Tiv3/gocmitra/core/util"
)

// OpenAIClient implements the LLMClient interface for OpenAI APIs.
type OpenAIClient struct{}

// Generate sends a prompt to the OpenAI API and returns a formatted commit message.
func (o OpenAIClient) Generate(prompt string, cfg config.Config) (string, error) {
	if strings.TrimSpace(prompt) == "" {
		logger.Error("OpenAI prompt is empty")
		return "", errors.New("OpenAI prompt is empty")
	}

	if strings.TrimSpace(cfg.APIKey) == "" {
		logger.Error("OpenAI API key is missing") // ✅ redacted, no exposure
		return "", errors.New("OpenAI API key is missing: ensure it is set in the environment and config is loaded correctly")
	}

	body := map[string]interface{}{
		"model": cfg.Model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	headers := map[string]string{
		"Authorization": "Bearer " + cfg.APIKey,
	}

	raw, err := sendGenerateRequest(cfg.Endpoint, body, headers)
	if err != nil {
		logger.Error(fmt.Sprintf("OpenAI request failed to %s: %v", util.Redact(cfg.Endpoint), err)) // ✅ redacted
		return "", fmt.Errorf("OpenAI request failed: %w", err)
	}

	var r struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		logger.Error(fmt.Sprintf("Failed to decode OpenAI response: %v\nRaw: %s", err, raw))
		return "", fmt.Errorf("OpenAI decode error: %w\nRaw: %s", err, raw)
	}

	if len(r.Choices) == 0 {
		logger.Warn("OpenAI returned no choices")
		return "", errors.New("OpenAI returned no choices")
	}

	content := strings.TrimSpace(r.Choices[0].Message.Content)
	if content == "" {
		logger.Warn("OpenAI returned empty content in the first choice")
		return "", errors.New("OpenAI returned empty content in the first choice")
	}

	content = strings.Trim(content, "`\".")
	return content, nil
}

// Name returns the name of this LLM provider.
func (o OpenAIClient) Name() string {
	return "OpenAI"
}
