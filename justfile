set shell := ["/bin/bash", "-c"]

import 'version.just'

default: help

help:
    @echo "Available recipes:"
    @just --list

# Install dependencies and setup development environment
[group("setup")]
install:
    go mod download
    go mod tidy

# Run tests
[group("test")]
test:
    go test -count=1 ./...

# Run tests with verbose output
[group("test")]
test-verbose:
    go test  -count=1 -v ./...

# Run tests with coverage
[group("test")]
test-coverage:
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html
    @echo "Coverage report generated: coverage.html"

# Run tests with integration tag
[group("test")]
test-integration:
    go test -tags=integration ./...

# Run integration tests with coverage  
[group("test")]
test-integration-coverage:
    go test -coverprofile=coverage-integration.out -tags=integration ./...
    go tool cover -html=coverage-integration.out -o coverage-integration.html
    @echo "Integration coverage report generated: coverage-integration.html"

# Run static analysis
[group("code")]
vet:
    go vet ./...

# Format code
[group("code")]
fmt:
    go fmt ./...

# Run golangci-lint
[group("code")]
lint:
    golangci-lint run

# Fix linting issues automatically
[group("code")]
lint-fix:
    golangci-lint run --fix


# Build the project
[group("build")]
build:
    go build ./...

# Clean build artifacts
[group("build")]
clean:
    go clean ./...
    rm -f coverage*.out coverage*.html
    rm -rf dist/

# Tidy go modules
[group("build")]
tidy:
    go mod tidy
