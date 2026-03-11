APP_NAME = mkdocs
CMD_PATH = ./cmd/cli
DIST_DIR = dist

.PHONY: install build uninstall deps

build:
	@go build -o $(APP_NAME) $(CMD_PATH)
	
install:
	@go build -o $(shell go env GOPATH)/bin/$(APP_NAME) $(CMD_PATH)

uninstall:
	@rm $(shell go env GOPATH)/bin/$(APP_NAME)

deps:
	@go mod download