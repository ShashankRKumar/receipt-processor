# Stage 1: Build
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o receipt-processor .

# Stage 2: Run
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

WORKDIR /app
COPY --from=builder /app/receipt-processor .

EXPOSE 8080
CMD ["./receipt-processor"]
