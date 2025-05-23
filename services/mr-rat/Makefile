.PHONY: all build build_implant build_listener clean

# Default operating system and architecture
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Output directory
BINDIR := bin

# Run the app
run:
	go run cmd/implant/main.go

# Build the implant binary
build:
	@echo "Building implant binary for GOOS=$(GOOS) GOARCH=$(GOARCH)..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINDIR)/implant ./cmd/implant

# Clean up binaries
clean:
	@echo "Cleaning up binaries..."
	rm -rf $(BINDIR)
