cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

.PHONY: help

help: ## Show help menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build an image from a Dockerfile with tag APP_NAME:VERSION
	docker build -t $(APP_NAME):$(VERSION) .

build-ne: ## Build the image if it does not exist
	if [ -z "$(shell docker images -q $(APP_NAME):$(VERSION))" ]; then docker build -t $(APP_NAME):$(VERSION) .; fi

rm: ## Delete the image
	docker image rm $(APP_NAME):$(VERSION) -f

run: build-ne ## Build the image if not exists and start containers
	docker run -it --rm -p=$(HOST_PORT):$(APP_PORT) --name="$(APP_NAME)" $(APP_NAME):$(VERSION)

up: build-ne ## Build the image if not exist and start containers using docker-compose
	docker-compose up --detach $(APP_NAME)

down: ## Stop and remove containers and networks created by docker-compose
	docker-compose down --rmi all

benchmark: up ## Run the containers and perform benchmark using apachebench
	docker run --rm -it --add-host host.docker.internal:host-gateway rafikurnia/ab:2.3-alpine3.15.4 -n $(NUMBER_OF_REQUESTS) -c $(NUMBER_OF_CONCURRENT_REQUESTS) -s 120 -k http://host.docker.internal/
	docker-compose down --rmi all

test: ## Test Go code
	go test -v -cover -count=1 ./...
