.env:
	cp .env.sample .env

.PHONY: prepare
prepare:
	cd backend && go mod vendor
	cd benchmark-service && go mod vendor

.PHONY: up
up: prepare .env
	docker compose up --wait --build
