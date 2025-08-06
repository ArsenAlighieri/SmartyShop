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

// GeminiAPIResponse is the raw response from the Google Gemini API.
type GeminiAPIResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// GeminiProductResponse is the structured JSON we expect Gemini to return.
type GeminiProductResponse struct {
	Answer   string             `json:"answer"`
	Products []internal.Product `json:"products"`
}

// GetGeminiProductInsights sends the product list and user question to Gemini API.
func GetGeminiProductInsights(products []internal.Product, userQuestion, apiKey string) (*GeminiProductResponse, error) {
	var productList strings.Builder
	for _, p := range products {
		productList.WriteString(fmt.Sprintf(
			"- Title: %s, Price: %s, Rating: %.1f, Reviews: %d, URL: %s\n",
			p.Title, p.Price, p.Rating, p.ReviewsCount, p.URL))
	}

	prompt := fmt.Sprintf(
		"You are an expert shopping assistant and a knowledgeable AI. Your primary goal is to provide comprehensive and detailed answers to user questions about products.\n\n"+
			"If the user asks for a recommendation, selection, or ranking, you must select up to 50 best products from the list based on their title, rating, and reviews count.\n"+
			"Crucially, if the question is not directly answerable from the provided products, you MUST use your extensive general knowledge and research capabilities to provide a comprehensive answer. Never state that you don't have enough information or that you can only answer based on provided products. Always strive to provide a detailed and informative response, even if it means drawing from your own knowledge base or simulating a web search.\n"+
			"If no relevant products are found or provided in the initial list, the 'products' array in the JSON can be empty, but you must still provide a detailed 'answer'.\n"+
			"When comparing products, provide detailed specifications and differences, similar to a product review site.\n"+
			"You must return your response as a single, raw JSON object and nothing else. Do not wrap it in markdown (e.g., ```json`). The JSON object must have the following structure:\n"+
			"{\n"+
			"  \"answer\": \"Your detailed answer to the user's question, summarizing your findings or recommendations. This answer should be rich in detail and comprehensive. Do not use any markdown formatting (e.g., *, **, #) within this answer field. Provide comparisons and specifications as if you have access to a vast product database.\",\n"+
			"  \"products\": [\n"+
			"    {\n"+
			"      \"title\": \"Product Title\",\n"+
			"      \"price\": \"Product Price\",\n"+
			"      \"rating\": 4.5,\n"+
			"      \"reviews_count\": 120,\n"+
			"      \"url\": \"Product URL\",\n"+
			"      \"image_url\": \"Image URL\",\n"+
			"      \"description\": \"Product Description\",\n"+
			"      \"site\": \"Site Name\"\n"+
			"    }\n"+
			"  ]\n"+
			"}\n\n"+
			"User question: %s\n\nProducts:\n%s",
		userQuestion,
		productList.String(),
	)

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash-lite:generateContent?key=%s", apiKey)

	requestBody, err := json.Marshal(GeminiRequest{
		Contents: []Content{
			{
				Parts: []Part{{Text: prompt}},
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

	// Remove Markdown formatting if Gemini returned it anyway
	jsonString := strings.TrimSpace(geminiAPIResp.Candidates[0].Content.Parts[0].Text)
	jsonString = strings.TrimPrefix(jsonString, "```json")
	jsonString = strings.TrimSuffix(jsonString, "```")
	jsonString = strings.TrimSpace(jsonString)

	var productResponse GeminiProductResponse
	if err := json.Unmarshal([]byte(jsonString), &productResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling product response JSON from Gemini: %w. Raw response: %s", err, jsonString)
	}

	return &productResponse, nil
}
