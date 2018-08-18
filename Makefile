EXECUTABLE=gameoflife
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

.PHONY: all test clean

all: test build ## Build and run tests

test: ## Run unit tests
	./scripts/test_unit.sh

bench: ## Run benchmark tests
	./scripts/test_bench.sh

dep: ## Retrieve all dependencies
	dep ensure -v

build: $(WINDOWS) ## Build binaries
	@echo version: $(VERSION)

$(WINDOWS): ## Build for Windows
	env GOOS=windows GOARCH=amd64 go build -i -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/game/main.go

$(DARWIN): ## Build for macOS
	env GOOS=darwin GOARCH=amd64 go build -i -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/game/main.go

$(LINUX): ## Build for Linux
	env GOOS=linux GOARCH=amd64 go build -i -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/game/main.go

clean: ## Remove previous build
	rm $(WINDOWS) $(LINUX) $(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'