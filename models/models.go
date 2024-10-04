package models

type AnthropicRequest struct {
	Model       string                    `json:"model"`
	MaxTokens   int                       `json:"max_tokens"`
	System      string                    `json:"system"`
	Temperature float32                   `json:"temperature"`
	Messages    []AnthropicRequestMessage `json:"messages"`
}

type AnthropicRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AnthropicResponseContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type AnthropicResponse struct {
	Id      string                     `json:"id"`
	Type    string                     `json:"type"`
	Role    string                     `json:"role"`
	Content []AnthropicResponseContent `json:"content"`
}
