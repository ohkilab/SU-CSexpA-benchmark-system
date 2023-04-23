.PHONY: prepare
prepare:
	cd backend && go mod vendor

.PHONY: build
build: prepare
	docker compose build

up:
	docker compose up --wait --build
