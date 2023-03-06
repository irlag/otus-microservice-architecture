include .env

GOLANG_VERSION=1.19
VERSION ?= $(shell git describe --tags 2> /dev/null || git rev-parse --short HEAD)
NAMESPACE=irlag
APP=otus-microservice-architecture

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
    		--tag ${NAMESPACE}/${APP}:${VERSION} \
    		--tag ${NAMESPACE}/${APP}:latest \
    		--file Dockerfile \
    		.

.PHONY: up
up: ## Up server
	GOLANG_VERSION=${GOLANG_VERSION} docker-compose up -d app

.PHONY: local-up
local-up: ## Up local server
	GOLANG_VERSION=${GOLANG_VERSION} docker-compose --profile dependencies up -d

.PHONY: exec
exec: ## Exec in app
	docker-compose exec app sh

.PHONY: run
run: ## Run app
	GOLANG_VERSION=${GOLANG_VERSION} docker-compose run --service-ports --rm app

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
	docker push ${NAMESPACE}/${APP}:${VERSION}
	docker push ${NAMESPACE}/${APP}:latest

.PHONY: sqlc-generate
sqlc-generate: ## Sqlc generate store
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate

#
# Deploy targets
#
.PHONY: create-ns
create-ns:
	kubectl create namespace ${NAMESPACE}
	kubectl label ns ${NAMESPACE} istio-injection=enabled

.PHONY: port-forward
port-forward:
	kubectl port-forward svc/istio-ingressgateway -n istio-system 8080:80

.PHONY: helm-install
helm-install: ## Deploy application using helm
	helm install otus-microservice-architecture ./ --namespace ${NAMESPACE} --create-namespace

.PHONY: helm-upgrade
helm-upgrade: ## Upgrade application using helm
	helm upgrade otus-microservice-architecture ./ --namespace ${NAMESPACE} --create-namespace

.PHONY: helm-uninstall
helm-uninstall: ## Uninstall application using helm
	helm uninstall otus-microservice-architecture --namespace ${NAMESPACE}

#
# Migrations targets
#
dev-migrate-create:
ifdef name
	docker-compose --profile dependencies run --rm migrate create -ext sql -dir /migrations/ $(name)
else
	$(error Error. Set migration name. Example $$ make dev-migrate-create name=init)
endif

dev-migrate-down dev-migrate-up: POSTGRESQL_URL= "postgres://$(APP):$(APP)@db:5432/$(APP)?sslmode=disable"
dev-migrate-up:
	$(info Start up migrations.)
	docker-compose --profile dependencies run --rm --service-ports migrate -database $(POSTGRESQL_URL) -path /migrations/ up

dev-migrate-down:
	$(info Start down migrations. Example $$ make dev-migrate-down count=1)
	docker-compose  --profile dependencies run --rm --service-ports migrate -database $(POSTGRESQL_URL) -path /migrations/ down $(count)
