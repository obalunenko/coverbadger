NAME=coverbadger

VERSION ?= $(shell git describe --tags $(git rev-list --tags --max-count=1))


# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=20


define colored
	@echo '${GREEN}$1${RESET}'
endef

## Show help
help:
	${call colored, help is running...}
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

build:
	./scripts/compile.sh
.PHONY: build

run: test-cover build
	./scripts/run.sh
.PHONY: run

test:
	./scripts/run-tests.sh
.PHONY: test

test-cover:
	./scripts/coverage.sh
.PHONY: cover

coverage:
	make cover

configure: sync-vendor

sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Fix imports sorting.
imports:
	${call colored, fix-imports is running...}
	./scripts/fix-imports.sh
.PHONY: imports

## Format code with go fmt.
fmt:
	${call colored, fmt is running...}
	./scripts/fmt.sh
.PHONY: fmt

## Format code and sort imports.
format-project: fmt imports
.PHONY: format-project

install-tools:
	./scripts/install-vendored-tools.sh
.PHONY: install-tools

## vet project
vet:
	${call colored, vet is running...}
	./scripts/run-vet.sh
.PHONY: vet

## Run full linting
lint-full:
	./scripts/run-linters.sh
.PHONY: lint-full

## Run linting for build pipeline
lint-pipeline:
	./scripts/run-linters-pipeline.sh
.PHONY: lint-pipeline

## recreate all generated code and swagger documentation.
codegen:
	${call colored, codegen is running...}
	./scripts/go-generate.sh
.PHONY: codegen

## recreate all generated code and swagger documentation and format code.
generate: codegen format-project vet
.PHONY: generate

## Release
release:
	./scripts/release.sh
.PHONY: release

## Release local snapshot
release-local-snapshot:
	${call colored, release is running...}
	./scripts/local-snapshot-release.sh
.PHONY: release-local-snapshot

## Issue new release.
new-version: vet test compile
	./scripts/new-version.sh
.PHONY: new-release