APP_NAME = linknet-smg-api
CURRENT_DIR := $(shell pwd)
BUILD_DIR = $(CURRENT_DIR)/build

build-linux:
	@echo "Building $(APP_NAME)..."
	GOOS=linux GOARCH=amd64 go build -C src -o $(BUILD_DIR)/$(APP_NAME)

build-win:
	@echo "Building for windows $(APP_NAME)..."
	GOOS=windows GOARCH=amd64 go build -C src -o $(BUILD_DIR)/$(APP_NAME).exe

clean:
	@echo "Cleaning $(APP_NAME)..."
	rm -rf $(BUILD_DIR)

run:
	@echo "Running $(APP_NAME)..."
	go run src/main.go


