# Receipt Processing API

A RESTful API service that processes retail receipts and calculates reward points based on specific business rules. The service provides a simple way to submit receipts and retrieve calculated points through HTTP endpoints.

## Table of Contents

- [Receipt Processing API](#receipt-processing-api)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Getting Started](#getting-started)
  - [API Documentation](#api-documentation)
    - [Endpoints](#endpoints)
      - [Post Receipts](#post-receipts)
      - [Get Points](#get-points)
  - [Future Improvements](#future-improvements)
  - [Author](#author)


## Overview

This API allows retailers to submit receipts and receive points based on specific rules like:

- One point for every alphanumeric character in the retailer name
- Points based on round dollar amounts and multiples of 0.25
- Points for even transaction dates
- Points based on purchase time (2:00 PM - 4:00 PM earns 10 points)
- Points for item descriptions and prices
- Points for receipt total and item count

## Features

✅ RESTful API endpoints for receipt processing and points retrieval
✅ In-memory storage with thread-safe operations using mutex locks
✅ Comprehensive request validation and error handling
✅ Environment-based configuration management
✅ Clean architecture following domain-driven design principles
✅ Separation of concerns with clear layer boundaries
✅ Type-safe request/response handling

## Requirements

- Go
- Git
- Air (Optional for hot reloading)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/geardestroyer85/receipt-api.git
   cd receipt-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure environment:
   - Copy `.env.example` to `.env`
   - Modify settings as needed

4. Run the application:
   
   Development mode with hot reload:
   ```bash
   air
   ```
   
   Standard development mode:
   ```bash
   go run src/cmd/api/main.go
   ```

   Production mode:
   ```bash
   go build -o build/prod/main.exe src/cmd/api/main.go
   ./build/prod/main.exe
   ```

## API Documentation

### Endpoints

#### Post Receipts

    POST /receipts/process  
    Content-Type: application/json

```json
{
  "retailer": "Target",
  "purchaseDate": "2024-01-01",
  "purchaseAmount": 100.00,
  "total": "100.00",
  "items": [
    {
      "shortDescription": "Item 1",
      "price": 50.00
    },
    {
      "shortDescription": "Item 2",
      "price": 50.00
    }
  ]
}
```

#### Get Points

    GET /receipts/{id}/points

## Future Improvements

- Add Swagger documentation
- Add logging system
- Add tests
- Add more error handling
- Add more validation
- Add more tests

## Author

- [@geardestroyer85](https://github.com/geardestroyer85)
