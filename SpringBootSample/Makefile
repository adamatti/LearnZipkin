.DEFAULT_GOAL := help

.PHONY: docker-build
docker-build:
	docker build -t spring-boot-sample:latest .

docker-sh: docker-build
	docker run -it --rm spring-boot-sample:latest sh

docker-run: docker-build
	docker run -it --rm spring-boot-sample:latest

.PHONY: help
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

