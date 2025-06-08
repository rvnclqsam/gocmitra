package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
)

// sendGenerateRequest is used for LLMs expecting a /generate-style prompt (e.g., OpenAI, Claude legacy).
func sendGenerateRequest(endpoint string, body map[string]interface{}, headers map[string]string) (string, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to marshal request body: %v", err))
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create HTTP request: %v", err))
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(fmt.Sprintf("HTTP request failed: %v", err))
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to read response body: %v", err))
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		logger.Warn(fmt.Sprintf("LLM error [%d]: %s", resp.StatusCode, string(bodyBytes)))
		return "", fmt.Errorf("LLM error [%d]: %s", resp.StatusCode, string(bodyBytes))
	}

	return string(bodyBytes), nil
}

// sendChatRequest handles /chat-style APIs and includes fallback decoding for Claude's /v1/messages response format.
func sendChatRequest(endpoint string, body map[string]interface{}, headers map[string]string) (string, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to marshal request body: %v", err))
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create HTTP request: %v", err))
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(fmt.Sprintf("HTTP request failed: %v", err))
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle Ollama separately (assumes streaming line-by-line JSON chunks)
	if strings.Contains(strings.ToLower(endpoint), "localhost") {
		var resultBuilder strings.Builder
		decoder := json.NewDecoder(resp.Body)

		for decoder.More() {
			var line struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			}
			if err := decoder.Decode(&line); err != nil {
				logger.Warn(fmt.Sprintf("Skipping malformed JSON chunk from Ollama: %v", err))
				continue
			}
			resultBuilder.WriteString(line.Message.Content)
		}

		result := strings.TrimSpace(resultBuilder.String())
		if result == "" {
			logger.Warn("Ollama stream returned no usable content")
			return "", fmt.Errorf("Ollama stream returned no usable content")
		}
		return result, nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to read response body: %v", err))
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		logger.Warn(fmt.Sprintf("LLM error [%d]: %s", resp.StatusCode, string(bodyBytes)))
		return "", fmt.Errorf("LLM error [%d]: %s", resp.StatusCode, string(bodyBytes))
	}

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var parsed struct {
			Content []struct {
				Text string `json:"text"`
			} `json:"content"`
		}
		if err := json.Unmarshal(bodyBytes, &parsed); err == nil && len(parsed.Content) > 0 {
			return strings.TrimSpace(parsed.Content[0].Text), nil
		}
	}

	return strings.TrimSpace(string(bodyBytes)), nil
}
