backend/.env:
	cp backend/.env.example backend/.env

.env:
	cp .env.example .env

.PHONY: prepare
prepare:
	cd backend && go mod vendor
	cd benchmark-service && go mod vendor

.PHONY: build
build: prepare
	docker compose build

.PHONY: up prepare
up:
	docker compose up --wait --build
