# Define the binary name
BINARY_NAME=server

# Define the source files
SOURCES=$(wildcard *.go)

# Default target
all: build

# Build the binary
build:
	@echo "Building the binary..."
	@go build -o $(BINARY_NAME) $(SOURCES)
	@echo "Build complete!"

# Run the binary
run: build
	@echo "Running the binary..."
	@./$(BINARY_NAME)

# Clean the build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "Clean complete!"

# PHONY targets
.PHONY: all build run clean
