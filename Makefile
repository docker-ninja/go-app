.PHONY: help
help: ## Prints help for all make commands
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

APP_NAME?=go-app

ensure-build-dir:
	mkdir -p out

build-deps: ## Install dependencies
	go install -mod=vendor
	go mod tidy
	go mod vendor

build: ensure-build-dir compile ## Build the application for mac

build-linux: ensure-build-dir compile-linux ## Build the application for linux

compile: ## Compile for mac
	go build -mod=vendor -o ./out/$(APP_NAME) ./main.go

compile-linux: ensure-build-dir ## Compile go-app for linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./out/$(APP_NAME) ./main.go
