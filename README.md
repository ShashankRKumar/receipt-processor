# 🧾 Receipt Processor API

The **Receipt Processor API** is a receipt processing service that enables users to submit receipt data (via a simple API request), extract important details, and calculate points based on predefined rules. The API is designed to be fast, scalable, and easy to use for developers looking to integrate receipt data processing into their systems.

## 🌟 Features

- **Receipt Parsing**: Extracts key details like retailer, purchase date, purchase time, items, and total amount from receipt data.
- **Points Calculation**: Calculates points based on the receipt details, such as items purchased and total amount.
- **Unique Receipt ID**: Each successfully processed receipt generates a unique ID.
- **Data Validation**: Ensures that receipt data is valid before processing.
- **Export Options**: Provides options to export the processed data in various formats (JSON, CSV).

---

## 🔧 Tech Stack

This API is built with the following technologies:

- **Go 1.21**: The backend of the application is built using Go, a fast and efficient programming language known for its performance and simplicity.
- **Gin Web Framework**: A web framework for Go that provides fast and flexible routing for HTTP requests.
- **Docker**: Used for containerizing the application, making it easier to deploy across different environments without worrying about dependency conflicts.
- **GORM**: The Object Relational Mapper (ORM) used for interacting with databases to store receipt information.

---

## 📦 Endpoints

### `POST /receipts/process`

This endpoint allows you to submit a receipt, and the server will process it and return a unique receipt ID.

#### **Description**
Submit a receipt to be processed. The API extracts key details from the receipt, such as retailer, purchase date, items, total amount, etc. Additionally, points will be calculated based on predefined rules.

#### **Request Payload**

The request should contain the following fields:

- `retailer`: The name of the retailer where the receipt was generated.
- `purchaseDate`: The date of the purchase in `YYYY-MM-DD` format.
- `purchaseTime`: The time of purchase in `HH:MM` format.
- `items`: A list of items purchased, including a short description of each item and the price.
- `total`: The total amount of the receipt.

**Example Request Payload**:

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
