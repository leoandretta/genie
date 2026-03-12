APP_NAME = mkdocs
CMD_PATH = ./cmd/cli
DIST_DIR = dist

ifeq ($(OS), Windows_NT)
    BINARY = $(APP_NAME).exe
    BIN_PATH = $(shell go env GOPATH)\\bin\\$(BINARY)
else
    BINARY = $(APP_NAME)
    BIN_PATH = $(shell go env GOPATH)/bin/$(BINARY)
endif

.PHONY: install build uninstall deps

build:
	@go build -o $(BINARY) $(CMD_PATH)
	
install:
	@go build -o $(BIN_PATH) $(CMD_PATH)

uninstall:	
	@rm $(BIN_PATH)

deps:
	@go mod download
