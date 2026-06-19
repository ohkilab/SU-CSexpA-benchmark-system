V2026_BUNDLE_DIR_OVERRIDE := $(V2026_BUNDLE_DIR)
V2026_CONTEST_SLUG_OVERRIDE := $(V2026_CONTEST_SLUG)
-include .env
ifneq ($(V2026_BUNDLE_DIR_OVERRIDE),)
V2026_BUNDLE_DIR := $(V2026_BUNDLE_DIR_OVERRIDE)
endif
ifneq ($(V2026_CONTEST_SLUG_OVERRIDE),)
V2026_CONTEST_SLUG := $(V2026_CONTEST_SLUG_OVERRIDE)
endif

.env:
	cp .env.sample .env

COMPOSE_BASE = docker compose -f compose.yaml -f compose.override.yaml
V2026_CONTEST_SLUG ?= v2026

.PHONY: prepare
prepare:
	@cd backend && go mod vendor
	@cd benchmark-service && go mod vendor
	@cd benchmark-service && make prepare

.PHONY: up
up: prepare .env
	$(COMPOSE_BASE) -f compose.sample.yaml up --wait --build

.PHONY: validate-v2026-bundle
validate-v2026-bundle:
	@test -n "$(V2026_BUNDLE_DIR)" || (echo "V2026_BUNDLE_DIR is required"; exit 1)
	@test -f "$(V2026_BUNDLE_DIR)/v2026.json" || (echo "$(V2026_BUNDLE_DIR)/v2026.json is required"; exit 1)
	@test -d "$(V2026_BUNDLE_DIR)/tags" || (echo "$(V2026_BUNDLE_DIR)/tags is required"; exit 1)
	@test -f "$(V2026_BUNDLE_DIR)/manifest.json" || (echo "$(V2026_BUNDLE_DIR)/manifest.json is required"; exit 1)

.PHONY: up-official
up-official: prepare .env validate-v2026-bundle
	V2026_BUNDLE_DIR="$(V2026_BUNDLE_DIR)" V2026_CONTEST_SLUG="$(V2026_CONTEST_SLUG)" $(COMPOSE_BASE) -f compose.official.yaml up --wait --build

.PHONY: clean
clean:
	@docker compose down --rmi local -v --remove-orphans
	@rm -rf ./data
