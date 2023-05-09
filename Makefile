tidy":
	go mod tidy

	# Define variables
APP_NAME=myapp
SRC=./cmd/main.go
BUILD_DIR=./bin
OUTPUT=$(BUILD_DIR)/$(APP_NAME)

# Build the program
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(OUTPUT) $(SRC)

# Clean up the build directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Run the program
run: build
	@echo "Running $(APP_NAME)..."
	@$(OUTPUT)

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Default target
.DEFAULT_GOAL := run
