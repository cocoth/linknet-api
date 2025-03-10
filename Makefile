APP_NAME = linknet-smg-api
BUILD_DIR = ./bin


build-linux:
	@echo "Building $(APP_NAME)..."
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)

build-win:
	@echo "Building for windows $(APP_NAME)..."
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME).exe

run:
	@echo "Running $(APP_NAME)..."
	go run src/main.go


