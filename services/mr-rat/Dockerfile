# Stage 1: Build the Go binary
FROM golang:1.18 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary, specifying the path to the main.go file
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mr-rat ./cmd/mr-rat

# Stage 2: Create a lightweight image to run the Go binary
FROM alpine:latest

# Install necessary CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/mr-rat .

# Set the environment variable
ENV APP_ENV=docker

# Expose the port the app runs on
EXPOSE 4002

# Command to run the Go binary
CMD ["./mr-rat"]
