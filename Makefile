.PHONY: build run clean docker-build docker-run docker-stop test fmt vet

# Build the application
build:
	@echo "Building TGBot4SillyTavern..."
	@go build -o tgbot4sillytavern .

# Run the application locally
run:
	@echo "Running TGBot4SillyTavern..."
	@go run main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f tgbot4sillytavern
	@go clean

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker-compose build

# Run with Docker Compose
docker-run:
	@echo "Starting with Docker Compose..."
	@docker-compose up -d
	@echo "Bot started! Check logs with: make docker-logs"

# Stop Docker containers
docker-stop:
	@echo "Stopping Docker containers..."
	@docker-compose down

# View Docker logs
docker-logs:
	@docker-compose logs -f

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	@go mod tidy

# Install development tools
dev-tools:
	@echo "Installing development tools..."
	@go install golang.org/x/tools/cmd/goimports@latest

# Run all checks (fmt, vet)
check: fmt vet
	@echo "All checks passed!"

# Help
help:
	@echo "Available commands:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application locally"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run with Docker Compose"
	@echo "  make docker-stop  - Stop Docker containers"
	@echo "  make docker-logs  - View Docker logs"
	@echo "  make test         - Run tests"
	@echo "  make fmt          - Format code"
	@echo "  make vet          - Run go vet"
	@echo "  make deps         - Download dependencies"
	@echo "  make tidy         - Tidy dependencies"
	@echo "  make check        - Run fmt and vet"
	@echo "  make help         - Show this help message"
