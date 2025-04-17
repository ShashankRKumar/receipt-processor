# Stage 1: Build
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o receipt-processor .

# Stage 2: Run
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /app/receipt-processor .

EXPOSE 8080

ENTRYPOINT ["/receipt-processor"]
