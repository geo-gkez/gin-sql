# Makefile for the Banking API

# Variables
APP_NAME=banking-api
MAIN_PATH=./cmd/banking/main.go
DOCKER_COMPOSE_FILE=deployments/docker-compose.yml

.DEFAULT_GOAL := help

.PHONY: fmt vet build run clean infra infra-down help

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o $(APP_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

clean:
	rm -f $(APP_NAME)
	go clean -cache

infra:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

infra-down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v

help:
	@echo "Makefile for the Banking API"
	@echo ""
	@echo "Usage:"
	@echo "  make fmt          # Format the code"
	@echo "  make vet          # Run go vet"
	@echo "  make build        # Build the application"
	@echo "  make run          # Run the application"
	@echo "  make clean        # Clean up build artifacts"
	@echo "  make infra        # Start infrastructure with Docker Compose"
	@echo "  make infra-down   # Stop infrastructure and remove volumes"
	@echo "  make help         # Show this help message"