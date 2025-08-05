# SmartyShop Backend

This is the backend for the SmartyShop application.

## Features

- Scrape products from Trendyol, Teknosa, MediaMarkt, and Amazon.
- Get the top 10 products based on rating.
- Ask questions about products using Gemini AI.

## API Endpoints

- `GET /products?site=<site>&query=<query>`: Scrape products from a given site.
  - `site`: `trendyol`, `teknosa`, `mediamarkt`, or `amazon`.
  - `query`: The product to search for.
- `GET /products/top10?site=<site>&query=<query>`: Get the top 10 products from the cache.
- `POST /gemini/query`: Ask a question about a list of products.
  - Body: `{"query": "<your_question>", "products": [<list_of_products>]}`

## Setup

1.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

2.  **Set environment variables:**

    Create a `.env` file in the `backend` directory and add your Gemini API key:

    ```
    GEMINI_API_KEY=<your_api_key>
    ```

3.  **Run the application:**

    ```bash
    go run cmd/main.go
    ```
