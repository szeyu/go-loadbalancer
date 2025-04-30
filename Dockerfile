# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary and env file from builder
COPY --from=builder /app/main .
COPY .env .

# Set environment variables from .env file (optional, as we're copying the file)
ENV PORT=8080
ENV API_KEYS=key1,key2,key3,key4,key5

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]
