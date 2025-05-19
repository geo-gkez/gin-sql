# Makefile for the Banking API

# Variables
APP_NAME=banking-api
MAIN_PATH=./cmd/banking/main.go

.DEFAULT_GOAL := build

.PHONY: fmt vet build run

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
