# GEMINI.md: SmartyShop Project

## Project Overview

This repository contains the source code for **SmartyShop**, a smart shopping assistant. The project is a web application with a Go backend and a frontend (currently in development).

The backend is a Go application that provides a REST API for the following functionalities:

*   **Product Scraping:** It can scrape product information (name, price, rating, etc.) from multiple e-commerce websites like Trendyol, Teknosa, MediaMarkt, and Amazon.
*   **Product Comparison:** It allows comparing prices and features of products from different vendors.
*   **AI-Powered Insights:** It uses the Gemini API to provide smart insights and summaries about products.

The backend is built using the **Gin** web framework and uses the `colly` library for web scraping.

The `frontend` directory is currently empty, but it is intended to house the user interface of the application.

## Building and Running

To run the backend server, you need to have Go installed on your system.

1.  **Set up the Environment:**
    Create a `.env` file in the `backend` directory and add your Gemini API key:
    ```
    GEMINI_API_KEY=your_api_key_here
    ```

2.  **Run the Server:**
    Navigate to the `backend` directory and run the following command:
    ```bash
    go run cmd/main.go
    ```
    The server will start on `http://localhost:8080`.

## API Endpoints

The backend exposes the following API endpoints:

*   `GET /products?site=<site>&query=<query>`: Scrapes and returns a list of products from the specified site for the given query.
*   `GET /products/top10?site=<site>&query=<query>`: Returns the top 10 rated products from the cached results.
*   `POST /gemini/query`: Sends a query and a list of products to the Gemini API for analysis and returns the insights.

## Development Conventions

*   **Dependency Management:** The project uses Go modules for dependency management. The dependencies are listed in the `go.mod` file.
*   **Code Structure:** The backend code is organized into several packages:
    *   `api`: Contains the API handlers.
    *   `scrapers`: Implements the scraping logic for different e-commerce sites.
    *   `gemini`: Handles the interaction with the Gemini API.
    *   `internal`: Defines the internal data structures.
    *   `config`: Manages the application configuration.
*   **Error Handling:** The application uses Go's standard error handling mechanisms.
*   **Logging:** The application uses the standard `log` package for logging.
