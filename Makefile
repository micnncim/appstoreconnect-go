# -----------------------------------------------------------------------------
# Global

SHELL       := bash
.SHELLFLAGS := -eu -o pipefail -c

MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory

.DEFAULT_GOAL := help

# -----------------------------------------------------------------------------
# Go

GO      ?= go
GO_TEST ?= gotestsum --hide-summary=skipped

OWNER  := micnncim
REPO   := appstoreconnect-go
MODULE := github.com/$(OWNER)/$(REPO)

# -----------------------------------------------------------------------------
# OpenAPI

PKG := appstoreconnect

PKG_VERSION := 0.1.0

APPSTORECONNECT_API_VERSION ?= 3.0

OPENAPI_SPEC_URL := https://developer.apple.com/sample-code/app-store-connect/app-store-connect-openapi-specification.zip

OPENAPI_SPEC_DIR      := api/openapi
OPENAPI_SPEC_RAW_DIR  := $(OPENAPI_SPEC_DIR)/raw
OPENAPI_SPEC_TMP_DIR  := $(OPENAPI_SPEC_DIR)/tmp
OPENAPI_SPEC_FILE     := $(OPENAPI_SPEC_DIR)/app-store-connect_$(APPSTORECONNECT_API_VERSION).json
OPENAPI_SPEC_RAW_FILE := $(OPENAPI_SPEC_RAW_DIR)/app-store-connect_$(APPSTORECONNECT_API_VERSION).json
OPENAPI_SPEC_PATCH    := $(OPENAPI_SPEC_DIR)/patches/app-store-connect_$(APPSTORECONNECT_API_VERSION).json

# -----------------------------------------------------------------------------
# Defines

define target
@printf "+ $(patsubst ,$@,$(1))\\n" >&2
endef

# -----------------------------------------------------------------------------
# Targets

##@ development

.PHONY: test
test: ## Run tests.
	$(call target)
	@$(GO_TEST) ./...

.PHONY: fmt
fmt: ## Run formatters.
fmt: fmt/go fmt/nix
	$(call target)

.PHONY: fmt/go
fmt/go: ## Format Go code.
	$(call target)
	@goimports -w -local $(MODULE) .
	@gofumpt -w .

.PHONY: fmt/nix
fmt/nix: ## Format Nix code.
	$(call target)
	@nix fmt

.PHONY: lint
lint: ## Run linters.
lint: lint/golangci-lint
	$(call target)

.PHONY: lint/golangci-lint
lint/golangci-lint: ## Run golanci-lint.
	$(call target)
	@golangci-lint run ./...

##@ code generation

.PHONY: generate
generate: ## Generate Go code.
generate: generate/openapi
	$(call target)

.PHONY: generate/openapi
generate/openapi: ## Generate Go code from OpenAPI spec.
generate/openapi: clean/openapi patch/openapi
	$(call target)
	@openapi-generator-cli generate \
		-g go \
		-i $(OPENAPI_SPEC_FILE) \
		--git-user-id $(OWNER) \
		--git-repo-id $(REPO) \
		--additional-properties packageName=$(PKG) \
		--additional-properties packageVersion=$(PKG_VERSION) \
		--additional-properties isGoSubmodule=true \
		--additional-properties enumClassPrefix=true \
		-o $(PKG)
	$(MAKE) fmt/go

.PHONY: patch/openapi
patch/openapi: ## Patch OpenAPI spec.
patch/openapi: $(OPENAPI_SPEC_RAW_FILE) $(OPENAPI_SPEC_PATCH)
	$(call target)
	@jq -s '.[0] * .[1]' $(OPENAPI_SPEC_RAW_FILE) $(OPENAPI_SPEC_PATCH) > $(OPENAPI_SPEC_FILE)

.PHONY: download/openapi
download/openapi: ## Download OpenAPI spec.
	$(call target)
	@mkdir -p $(OPENAPI_SPEC_TMP_DIR)
	@curl -so $(OPENAPI_SPEC_TMP_DIR)/raw.zip $(OPENAPI_SPEC_URL)
	@unzip -q $(OPENAPI_SPEC_TMP_DIR)/raw.zip -d $(OPENAPI_SPEC_TMP_DIR)
	@jq . $(OPENAPI_SPEC_TMP_DIR)/openapi-deployment-publicApi-publicApi_$(APPSTORECONNECT_API_VERSION).json > $(OPENAPI_SPEC_RAW_FILE)
	@rm -rf $(OPENAPI_SPEC_TMP_DIR)

.PHONY: clean
clean: ## Clean files.
clean: clean/openapi
	$(call target)

.PHONY: clean/openapi
clean/openapi: ## Clean OpenAPI files.
	$(call target)
	@find $(PKG) -mindepth 1 ! -name .openapi-generator-ignore -exec rm -rf {} +

##@ help

.PHONY: help
help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[33m<target>\033[0m\n"} /^[a-zA-Z_0-9\/%_-]+:.*?##/ { printf "  \033[1;34m%-30s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
