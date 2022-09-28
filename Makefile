SHELL := /bin/bash

COVERAGE_REPORT_FILE := coverage.out

test: ## Run unit tests and generate report file
	go test -coverprofile=$(COVERAGE_REPORT_FILE) -v

coverage: ## Generate a graphical coverage report in HTML format based in report file
	@if [ -f "$(COVERAGE_REPORT_FILE)" ]; then \
        go tool cover -html=$(COVERAGE_REPORT_FILE) -o coverage.html; \
    else \
        make test && \
		go tool cover -html=$(COVERAGE_REPORT_FILE) -o coverage.html; \
    fi

help: ## Display help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.PHONY: help coverage test