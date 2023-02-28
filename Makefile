include .env

GOLANG_VERSION=1.19
VERSION ?= $(shell git describe --tags 2> /dev/null || git rev-parse --short HEAD)
NAMESPACE=irlag/otus-microservice-architecture

.DEFAULT_GOAL := help

.PHONY: help
help:	## Display this help
	@grep -hE '^[A-Za-z0-9_ \-]*?:.*##.*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## Display current version
	@echo "VERSION=${VERSION}"

.PHONY: build
build: ## Build app
	@docker build \
			--build-arg VERSION=${VERSION} \
    		--build-arg GOLANG_VERSION=${GOLANG_VERSION} \
    		--tag ${NAMESPACE}:${VERSION} \
    		--tag ${NAMESPACE}:latest \
    		--file Dockerfile \
    		.

.PHONY: up
up: ## Up server
	GOLANG_VERSION=${GOLANG_VERSION} docker-compose up -d app

.PHONY: exec
exec: ## Exec in app
	docker-compose exec app sh

.PHONY: run
run: ## Run app
	docker-compose run --service-ports --rm app

PHONY: stop
stop: ## Stop app
	docker-compose stop

PHONY: down
down: ## Down app
	docker-compose down --remove-orphans

.PHONY: remove
remove: ## Down and remove all containers
	docker-compose down -v --rmi all

.PHONY: push
push: ## Push image
	docker push ${NAMESPACE}:${VERSION}
	docker push ${NAMESPACE}:latest
