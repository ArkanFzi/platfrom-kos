# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install dependencies for build (if needed)
RUN apk add --no-cache git

# Copy gomod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final Stage
FROM alpine:latest

WORKDIR /app

# Install tzdata for timezone support
RUN apk add --no-cache tzdata

# Copy binary from builder
COPY --from=builder /app/main .

# Copy environment file (optional, or rely on docker-compose)
# COPY .env . 

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]
