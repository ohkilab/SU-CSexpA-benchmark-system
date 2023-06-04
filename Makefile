.env:
	cp .env.sample .env

.PHONY: prepare
prepare:
	@cd backend && go mod vendor
	@cd benchmark-service && go mod vendor
	@cd benchmark-service && make prepare

.PHONY: up
up: prepare .env
	docker compose up --wait --build
