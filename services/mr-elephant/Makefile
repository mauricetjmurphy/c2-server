MODULE_NAME := github.com/mauricetjmurphy/mr-elephant
APP_NAME    := mr-elephant
ECR_REPO    := your-account-id.dkr.ecr.your-region.amazonaws.com/$(APP_NAME)
ECR_REGION  := us-east-1

swag-fmt:
	swag fmt -g cmd/${APP_NAME}/main.go

swag: swag-fmt
	swag init -g cmd/${APP_NAME}/main.go --outputTypes go,json,yaml --parseDependency

# Format code
format:
	go fmt ./...

# Clean Go cache
clean-cache:
	go clean -cache -modcache -i -r

# Docker build
docker-build:
	docker build -t $(APP_NAME) .

# Docker build
docker-run: docker-build
	docker run -p 4000:4000 $(APP_NAME)

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

# Build the API locally
api-build:
	rm -f ./build/$(APP_NAME)
	go build -o ./build/$(APP_NAME) $(MODULE_NAME)/cmd/$(APP_NAME)

# Run the API locally
run: api-build
	./build/$(APP_NAME)

# Clean the built binary
clean:
	rm -f ./build/$(APP_NAME)

.PHONY: format clean-cache build docker-build docker-tag docker-push clean deploy
