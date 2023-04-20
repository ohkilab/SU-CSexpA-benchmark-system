.PHONY: build

backend/vendor:
	cd backend && go mod vendor

build: backend/vendor
	docker compose build

up:
	docker compose up --wait --build
