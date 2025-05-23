# Stage 1: Build the Go binary
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary, specifying the path to the main.go file
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mr-elephant ./cmd/mr-elephant

# Stage 2: Create a lightweight image to run the Go binary
FROM alpine:latest

# Install necessary CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/mr-elephant .

# Ensure the configuration directory exists and copy the config file
RUN mkdir -p /root/internal/api/config
COPY ./internal/api/config/local.yaml /root/internal/api/config/local.yaml

# Set the environment variables
ENV SERVER_ADDR=docker

# Expose the port the app runs on
EXPOSE 4000

# Command to run the Go binary
CMD ["./mr-elephant"]
