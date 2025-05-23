#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Define variables
APP_NAME="basic-api"
ECR_REGION="us-east-1"
ECR_REPO_NAME="c2-server"
ECR_REPO="144817152095.dkr.ecr.$ECR_REGION.amazonaws.com/$ECR_REPO_NAME"

echo "Tagging the Docker image..."
docker tag $APP_NAME:latest $ECR_REPO:latest

echo "Logging in to Amazon ECR..."
aws ecr get-login-password --region $ECR_REGION | docker login --username AWS --password-stdin $ECR_REPO

echo "Pushing the Docker image to ECR..."
docker push $ECR_REPO:latest

echo "Deploy process completed successfully."
