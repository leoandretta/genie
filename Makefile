APP_NAME = mkdocs
CMD_PATH = ./cmd/cli
DIST_DIR = dist

ifeq ($(OS), Windows_NT)
    BINARY = $(APP_NAME).exe
    BIN_PATH = $(shell go env GOPATH)\bin\$(BINARY)
		RM_CMD = del /f
else
    BINARY = $(APP_NAME)
    BIN_PATH = $(shell go env GOPATH)/bin/$(BINARY)
		RM_CMD = rm -f
endif

.PHONY: install build uninstall deps

build:
	@go build -o $(BINARY) $(CMD_PATH)
	
install:
	@go build -o $(BIN_PATH) $(CMD_PATH)
	@echo "$(APP_NAME) installed successfully"

uninstall:	
	@$(RM_CMD) $(BIN_PATH)
	@echo "$(APP_NAME) uninstalled successfully"

deps:
	@go mod download


