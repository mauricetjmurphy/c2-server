#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Define variables
APP_NAME="basic-api"

echo "Formatting the code..."
go fmt ./...

echo "Cleaning Go cache..."
go clean -cache -modcache -i -r

echo "Building the Docker image..."
docker build -t $APP_NAME .

echo "Build process completed successfully."
