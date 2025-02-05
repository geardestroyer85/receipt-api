# Receipt Processing API

A RESTful API service that processes receipts and calculates reward points based on specific rules.

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
