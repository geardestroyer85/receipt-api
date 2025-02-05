# Receipt Processing API

A RESTful API service that processes receipts and calculates reward points based on specific rules.

## Table of Contents

- [Receipt Processing API](#receipt-processing-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Usage](#usage)
    - [Development Mode](#development-mode)
    - [Production Mode](#production-mode)
  - [API Documentation](#api-documentation)
    - [API Endpoints](#api-endpoints)
      - [Post Receipts](#post-receipts)
      - [Get Points](#get-points)

## Features

- [ ] Process receipts and calculate points
- [ ] Retrieve points for processed receipts
- [ ] Configurable Swagger documentation
- [ ] Comprehensive logging system
- [ ] Environment configuration
- [ ] Error handling and validation

## Requirements

- Go
- Git

## Usage

### Development Mode

- Using air (live reload)  
  `air`

- Or using standard Go  
  `go run src/cmd/api/main.go`

### Production Mode

- Build the application  
  `go build -o build/prod/main.exe src/cmd/api/main.go`

- Run the binary  
  `./build/prod/main.exe`

## API Documentation

### API Endpoints

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
