SERVICE_NAME ?=$(notdir $(shell pwd))

.PHONY: all
all: help

.PHONY: fmt
fmt: ## コードをフォーマットします
	goimports -l -w .

.PHONY: tidy
tidy: ## Updates the go modules and vendors all dependencies
	go mod tidy
	go mod vendor

.PHONY: run
run: tidy ## Runs uncompiled code in Dapr
	dapr run \
		--app-id $(SERVICE_NAME) \
		--app-port 3002 \
		--app-protocol grpc \
		--dapr-http-port 3502 \
		--components-path ../../components \
		go run main.go

.PHONY: __
__:
	@echo "\033[33m"
	@echo "kzmake/dapr-clock"
	@echo "\033[0m"

.PHONY: help
help: __ ## shows help
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@cat $(MAKEFILE_LIST) \
	| grep -e "^[a-zA-Z_/\-]*: *.*## *" \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}'
