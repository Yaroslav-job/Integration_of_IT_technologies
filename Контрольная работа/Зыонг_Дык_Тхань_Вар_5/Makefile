# Makefile for Go Algorithms Project
# Student: Duong_Duc_Thanh
# License: GPLv3

# Variables
BINARY_NAME=goalgorithms
MAIN_PACKAGE=github.com/user/goalgorithms
GO=go
GOFLAGS=-ldflags="-s -w"
BUILD_DIR=./bin

.PHONY: all clear build run test deploy

# Default target
all: build

# Clean up build artifacts
clear:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	$(GO) clean
	@echo "Clean complete"

# Build the application
build:
	@echo "Building..."
	mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the application
run: build
	@echo "Running application..."
	$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...
	@echo "Tests complete"

# Deploy (for demonstration, simply copies the binary to a "deploy" directory)
deploy: build
	@echo "Deploying..."
	mkdir -p ./deploy
	cp $(BUILD_DIR)/$(BINARY_NAME) ./deploy/
	@echo "Deployment complete: ./deploy/$(BINARY_NAME)"