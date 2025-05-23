# Define variables
MODULE_NAME := github.com/mauricetjmurphy/service-01
APP_NAME := basic-api
ECR_REPO := your-account-id.dkr.ecr.your-region.amazonaws.com/$(APP_NAME)
ECR_REGION := us-east-1

# Format code
format:
	go fmt ./...

# Clean Go cache
clean-cache:
	go clean -cache -modcache -i -r

# Docker build
docker-build:
	docker build -t $(APP_NAME) .

# Tag the Docker image
docker-tag:
	docker tag $(APP_NAME):latest $(ECR_REPO):latest

# Push the Docker image to ECR
docker-push:
	aws ecr get-login-password --region $(ECR_REGION) | docker login --username AWS --password-stdin $(ECR_REPO)
	docker push $(ECR_REPO):latest

# Build the application using build.sh
build:
	./scripts/build.sh

# Deploy to ECR using deploy.sh
deploy:
	./scripts/deploy.sh

# Clean the built binary
clean:
	rm -f $(APP_NAME)

.PHONY: format clean-cache build docker-build docker-tag docker-push clean deploy
