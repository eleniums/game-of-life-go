EXECUTABLE=gameoflife.exe
VERSION=$(shell git describe --tags --always --long --dirty)

all: test build ## Build and run tests

test: ## Run unit tests
	./scripts/test_unit.sh

bench: ## Run benchmark tests
	./scripts/test_bench.sh

dep: ## Retrieve all dependencies
	dep ensure

build: $(EXECUTABLE) ## Build executable

$(EXECUTABLE):
	@echo version: $(VERSION)
	go build -i -v -o $(EXECUTABLE) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/game/main.go

clean: ## Remove previous build
	rm $(EXECUTABLE)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'