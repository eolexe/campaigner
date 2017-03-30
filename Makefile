PWD=$(shell pwd)
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

GIT_COMMIT = $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date -u +"%FT%T%z")
APP_VERSION = "${GIT_COMMIT}_${BUILD_DATE}"

BINARY = campaigner
LDFLAGS = -ldflags "-X main.Version=${APP_VERSION}"

.DEFAULT_GOAL=help

build: ## build the binary
	@echo "$(OK_COLOR)==> Building$(NO_COLOR)"
	go build ${LDFLAGS} -o ${BINARY} cmd/server/main.go

.PHONY: setup
setup: ## setup the local environment by installing required tools
	@echo "$(OK_COLOR)==> Setting up local environment$(NO_COLOR)"
	go get -u github.com/cespare/reflex

.PHONY: run.local
run.local: ## runs application locally
	@echo "$(OK_COLOR)==> Running campaigner service locally $(NO_COLOR)"
	ulimit -S -n 5000 && reflex -r '\.(go|json|yml)$$' -R '^vendor/' -s -- sh -c 'go build ${LDFLAGS} -o ${BINARY} cmd/server/main.go && ./${BINARY}'

.PHONY: run.tests
run.tests: ## runs tests
	@echo "$(OK_COLOR)==> Running tests $(NO_COLOR)"
	cd model && go test

.PHONY: run.benchmarks
run.benchmarks: ## runs benchmarks only (no tests!)
	@echo "$(OK_COLOR)==> Running benchmarks $(NO_COLOR)"
	cd model && go test -run=XXX -bench=.

.PHONY: docker.start
docker.start: ## build and start docker environment
	@echo "$(OK_COLOR)==> Start local docker environment $(NO_COLOR)"
	docker-compose up -d --build

.PHONY: docker.stop
docker.stop: ## stop docker environment
	@echo "$(OK_COLOR)==> Stop local docker environment $(NO_COLOR)"
	docker kill campaigner.redis

.PHONY: help
help:
	@grep -E '^[a-zA-Z\._-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
