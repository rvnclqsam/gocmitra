package llm

import (
	"encoding/json"
	"fmt"

	"github.com/Cre4T3Tiv3/gocmitra/core/config"
	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
)

type ClaudeClient struct{}

func (c ClaudeClient) Generate(prompt string, cfg config.Config) (string, error) {
	if cfg.APIKey == "" {
		logger.Error("Claude API key is missing")
		return "", fmt.Errorf("Claude API key is missing")
	}

	fullPrompt := fmt.Sprintf(
		"%s\n\nInstructions: You are an assistant that writes a single-line Git commit message in conventional format using a %s tone and %s style. "+
			"Do not include any explanation or bullet points. Output only the commit message.",
		prompt, cfg.Tone, cfg.Style,
	)

	body := map[string]interface{}{
		"model":       cfg.Model,
		"max_tokens":  300,
		"temperature": 0.5,
		"messages": []map[string]string{
			{"role": "user", "content": fullPrompt},
		},
	}

	headers := map[string]string{
		"x-api-key":         cfg.APIKey,
		"anthropic-version": "2023-06-01",
	}

	raw, err := sendGenerateRequest(cfg.Endpoint, body, headers)
	if err != nil {
		logger.Error(fmt.Sprintf("Claude request failed: %v", err))
		return "", err
	}

	var response struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}

	if err := json.Unmarshal([]byte(raw), &response); err != nil {
		logger.Error(fmt.Sprintf("Failed to decode Claude response: %v\nRaw: %s", err, raw))
		return "", err
	}

	for _, part := range response.Content {
		if part.Type == "text" && part.Text != "" {
			return part.Text, nil
		}
	}

	logger.Warn("Claude returned no usable text content")
	return "", fmt.Errorf("Claude returned empty content")
}

func (c ClaudeClient) Name() string {
	return "Claude"
}
