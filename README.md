#  Receipt Processor API

This project implements a receipt processing API that calculates points based on a set of rules.

##  Tech Stack

- Go 1.21
- Gin Web Framework
- Docker

---

##  Endpoints

### `POST /receipts/process`

**Description**: Submit a receipt and receive a unique ID.

**Example Payload**:

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    { "shortDescription": "Mountain Dew 12PK", "price": "6.49" },
    { "shortDescription": "Emils Cheese Pizza", "price": "12.25" }
  ],
  "total": "18.74"
}
