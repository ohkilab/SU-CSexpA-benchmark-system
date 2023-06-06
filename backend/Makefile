.PHONY: generate
generate:
	go generate ./server/repository/ent

.PHONY: mysql/start
mysql/start:
	$(eval CONTAINER_ID := $(shell docker run -d -P \
		-e MYSQL_DATABASE=su-exp-a \
		-e MYSQL_ROOT_PASSWORD=ohkilab \
		mysql:8.0))
	@echo $(CONTAINER_ID)

.PHONY: e2e_test
e2e_test: mysql/start
	@TEST_MYSQL_PORT=$(shell docker port $(CONTAINER_ID) | rg 3306/tcp | awk -F ' -> ' '{print $$2}' | awk -F ':' '{print $$2}') \
	MYSQL_ROOT_PASSWORD=ohkilab \
	MYSQL_DATABASE=su-exp-a \
		go test ./server/test/e2e -v
	@docker rm -f $(CONTAINER_ID)

.PHONY: batch/seed
batch/seed:
	@DB_USER=root \
	DB_PASS=root \
	DB_HOST=localhost \
	DB_PORT=3306 \
	DB_NAME=exp-a \
	go run ./cmd/batch seed

.PHONY: batch/create-users
batch/create-users:
	@DB_USER=root \
	DB_PASS=root \
	DB_HOST=localhost \
	DB_PORT=3306 \
	DB_NAME=exp-a \
	go run ./cmd/batch create-users .create-users.yaml

.PHONY: seed
seed: batch/seed
	@echo "seed is deprecated. please use \"make batch/seed\""

.PHONY: viz
viz:
	@open ./ent/schema-viz.html
