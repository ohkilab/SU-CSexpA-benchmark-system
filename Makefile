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

.PHONY: clean
clean:
	@docker compose down --rmi local -v --remove-orphans
	@rm -rf ./data
