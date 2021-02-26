.DEFAULT_GOAL := help

.PHONY: help
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## build docker images
	docker-compose build

up: build ## start full sample
	docker-compose up -d

start: up 

stop: ## stop containers
	docker-compose stop -t 0

restart: stop up ## restart sample

test-node:
	curl http://localhost:3000/people/1

test-spring:
	curl http://localhost:8080/people/1

test-go:
	curl http://localhost:8000/people/1

test: test-node test-spring test-go
	