REPO                  := github.com/PutskouDzmitry/golang-training-Library

PHONY: help
help: ## makefile targets description
	@echo "Usage:"
	@egrep '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##/#-/' | column -t -s "#"

.PHONY: fmt
fmt: ## automatically formats Go source code
	@echo "Running 'go fmt ...'"
	@go fmt -x "$(REPO)/..."

.PHONY: image
image: fmt ## build image from Dockerfile ./docker/server/Dockerfile
	@docker build -t kvarc/itest-app-v1 .


.PHONY: up
up : image ## up docker compose
	@docker-compose up -d

.PHONY: test
test: up
	@go test ./cmd/server/cmd_test.go

.PHONY: down
down : test
	@docker-compose down

