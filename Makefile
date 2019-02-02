.DEFAULT_GOAL := help

build: ## build docker images
	docker-compose build

up: build ## start full sample
	docker-compose up -d

stop: ## stop containers
	docker-compose stop -t 0

restart: stop up ## restart sample

.PHONY: help
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
