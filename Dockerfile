# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tgbot4sillytavern .

# Runtime stage
FROM chromedp/headless-shell:latest

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy binary from builder
COPY --from=builder /app/tgbot4sillytavern .

# Set environment variables
ENV HEADLESS_MODE=true
ENV DEBUG=false

# Run the application
CMD ["./tgbot4sillytavern"]
