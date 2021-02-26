.DEFAULT_GOAL := help

.PHONY: help
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install:
	go mod download	

build: ## build project
	go build -o sample

run: build ## run the 	binary
	./sample

docker-build:
	docker build . -t gosample:latest

docker-run: docker-build
	docker run --rm -p 8000:8000 gosample:latest