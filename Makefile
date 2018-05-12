EXECUTABLE=gameoflife.exe

all: build test ## Build and run tests

test: ## Run unit tests
	./scripts/test_unit.sh

bench: ## Run benchmark tests
	./scripts/test_bench.sh

dep: ## Retrieve all dependencies
	dep ensure

build: $(EXECUTABLE) ## Build executable

$(EXECUTABLE):
	go build -o $(EXECUTABLE) ./cmd/game/main.go

clean: ## Remove previous build
	rm $(EXECUTABLE)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'