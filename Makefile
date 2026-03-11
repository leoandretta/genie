APP_NAME = mkdocs
CMD_PATH = ./cmd/cli
DIST_DIR = dist

ifeq ($(OS), Windows_NT)
    BINARY = $(APP_NAME).exe
else
    BINARY = $(APP_NAME)
endif

.PHONY: install build uninstall deps

build:
	@go build -o $(BINARY) $(CMD_PATH)
	
install:
	@go build -o $(shell go env GOPATH)/bin/$(BINARY) $(CMD_PATH)

uninstall:.	
	@rm $(shell go env GOPATH)/bin/$(BINARY)

deps:
	@go mod download