package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"smartyshop/internal"
	"strings"
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

// GeminiAPIResponse is the raw response from the Google Gemini API
type GeminiAPIResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// GeminiProductResponse is the structured JSON we expect Gemini to return in the 'text' field.
type GeminiProductResponse struct {
	Answer   string             `json:"answer"`
	Products []internal.Product `json:"products"`
}

// GetGeminiProductInsights sends the product list and user question to the Gemini API and returns the structured response.
func GetGeminiProductInsights(products []internal.Product, userQuestion string, apiKey string) (*GeminiProductResponse, error) {
	var productList strings.Builder
	for _, p := range products {
		// Including more details for better analysis by the AI
		productList.WriteString(fmt.Sprintf("- Title: %s, Price: %s, Rating: %.1f, Reviews: %d, URL: %s\n",
			p.Title, p.Price, p.Rating, p.ReviewsCount, p.URL))
	}

	prompt := fmt.Sprintf(`You are an expert shopping assistant. Your task is to analyze the provided list of products and answer the user's question.

If the user asks for a recommendation, selection, or ranking, you must select up to 50 best products from the list based on their title, rating, and reviews count.

You must return your response as a single, raw JSON object and nothing else. Do not wrap it in markdown (e.g., `+"```json"+`). The JSON object must have the following structure:
{
  "answer": "Your detailed answer to the user's question, summarizing your findings or recommendations.",
  "products": [
    {
      "title": "Product Title",
      "price": "Product Price",
      "rating": 4.5,
      "reviews_count": 120,
      "url": "Product URL",
      "image_url": "Image URL",
      "description": "Product Description",
      "site": "Site Name"
    }
  ]
}

User question: %s

Products:
%s
`, userQuestion, productList.String())

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
		return nil, fmt.Errorf("error marshalling request: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error making request to Gemini API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Gemini API returned non-200 status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var geminiAPIResp GeminiAPIResponse
	if err := json.Unmarshal(body, &geminiAPIResp); err != nil {
		return nil, fmt.Errorf("error unmarshalling Gemini API response: %w. Body: %s", err, string(body))
	}

	if len(geminiAPIResp.Candidates) == 0 || len(geminiAPIResp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no content in Gemini response")
	}

	jsonString := geminiAPIResp.Candidates[0].Content.Parts[0].Text
	jsonString = strings.TrimPrefix(jsonString, "```json")
	jsonString = strings.TrimSuffix(jsonString, "```")
	jsonString = strings.TrimSpace(jsonString)

	var productResponse GeminiProductResponse
	if err := json.Unmarshal([]byte(jsonString), &productResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling product response JSON from Gemini: %w. Raw response from AI: %s", err, jsonString)
	}

	return &productResponse, nil
}
