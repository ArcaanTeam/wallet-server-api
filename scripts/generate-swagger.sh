#!/bin/bash

echo "Generating Swagger documentation..."

# Install swag if not exists
if ! command -v swag &> /dev/null; then
    echo "Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate swagger docs
swag init -g cmd/server/main.go --output docs --parseDependency --parseInternal

echo "Swagger documentation generated successfully!"