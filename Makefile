APP_NAME = linknet-api
CURRENT_DIR := $(shell pwd)
BUILD_DIR = $(CURRENT_DIR)/build

# Ensure BUILD_DIR exists
$(BUILD_DIR):
	@mkdir -p $(BUILD_DIR)

build-linux: $(BUILD_DIR)
	@echo "Cleaning $(APP_NAME)..."
	rm -rf $(BUILD_DIR)/*
	@echo "Building $(APP_NAME) for Linux..."
	cp .env $(BUILD_DIR)
	cp -r ./config/* $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -C src -o $(BUILD_DIR)/$(APP_NAME)

build-win: $(BUILD_DIR)
	@echo "Cleaning $(APP_NAME)..."
	rm -rf $(BUILD_DIR)/*
	@echo "Building $(APP_NAME) for Windows..."
	cp .env $(BUILD_DIR)
	cp -r ./config/* $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -C src -o $(BUILD_DIR)/$(APP_NAME).exe

clean:
	@echo "Cleaning $(APP_NAME)..."
	rm -rf $(BUILD_DIR)

run:
	@echo "Running $(APP_NAME)..."
	go run src/main.go serve

prod: $(BUILD_DIR)
	@echo "Running $(APP_NAME) in production mode..."
	export GIN_MODE=release && $(BUILD_DIR)/$(APP_NAME) serve