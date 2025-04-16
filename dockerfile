# Use the official Golang image as the base image
FROM golang:1.24.2 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the working directory
COPY . .

# Build the Go app inside the container
RUN go build -o main .

# Start a new stage to reduce image size (using a minimal base image)
FROM debian:bullseye-slim

# Install necessary dependencies to run Go binary
RUN apt-get update && apt-get install -y ca-certificates

# Set the working directory for the runtime environment
WORKDIR /app

# Copy the compiled binary from the builder stage to the runtime image
COPY --from=builder /app/main .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
