# StockApp Backend

This is the backend service for StockApp, a stock recommendation and management system written in Go. It provides APIs for fetching stock data, managing stock information, and generating stock recommendations.
This project was built for a hiring challenge, which included a provided data API and a set of requirements for the backend service.

## Features
- Fetches and stores stock data (From a provided API)
- Provides stock buying recommendations
- Supports querying stocks with various filters and sorting options
- Modular architecture (controller, service, repository, model)

## Project Structure
```
go.mod                      # Go module definition
go.sum                      # Go dependencies
main.go                     # Application entry point
controller/                 # HTTP controllers (API endpoints)
  stock-controller.go       # Stock-related endpoints
db/                         # Database connection and repository
  db.go                     # DB connection logic
  stock-repository.go       # Stock data repository
fetcher/                    # External data fetching logic
  fetcher.go                # Fetches stock data from APIs
model/                      # Data models and DTOs
  dto.go                    # Data Transfer Objects
  stock.go                  # Stock model
db/                         # Database logic
service/                    # Business logic
  recommendation-service.go # Stock recommendation logic
  stock-service.go          # Stock management logic
  weigths.go                # Recommendation weights
```

## Getting Started

### Prerequisites
- Go 1.18 or newer

### Installation
1. Clone the repository:
   ```sh
   git clone <repo-url>
   cd back
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

### Running the Application
```sh
go run main.go
```

## API Endpoints
- `/stock` - Get stock by ID
- `/stocks` - Get list of stocks
- `/sync` - Update stock data from external API
- `/recommendations` - Get stock recommendations
- `/query-stocks` - Query stocks filtered and sorted by various parameters
