package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"smartyshop/internal"
)

// GeminiRequest represents the request to the Gemini API.

type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

// Content represents the content of the request.

type Content struct {
	Parts []Part `json:"parts"`
}

// Part represents a part of the content.

type Part struct {
	Text string `json:"text"`
}

// GeminiResponse represents the response from the Gemini API.

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

// Candidate represents a candidate in the response.

type Candidate struct {
	Content Content `json:"content"`
}

// GetGeminiProductInsights sends the product list and user question to the Gemini API and returns the response.
func GetGeminiProductInsights(products []internal.Product, userQuestion string, apiKey string) (string, error) {
	prompt := fmt.Sprintf("User question: %s\n\nProducts:\n", userQuestion)
	for _, p := range products {
		prompt += fmt.Sprintf("- %s: %s\n", p.Title, p.Price)
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=%s", apiKey)

	requestBody, err := json.Marshal(GeminiRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: prompt},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var geminiResp GeminiResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		return "", err
	}

	if len(geminiResp.Candidates) > 0 {
		return geminiResp.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("no response from Gemini")
}
