.PHONY: build

backend/vendor:
	cd backend && go mod vendor

build: backend/vendor
	docker compose build
