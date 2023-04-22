ARCH=$(shell uname -m)
ifeq ($(ARCH),arm64)
	ARCH := aarch64
endif
OS=$(shell uname -s | tr A-Z a-z)
PROTOC_GEN_GRPC_WEB_VERSION=1.4.2

.PHONY: tools/install
tools/install:
	@mkdir -p bin \
		&& wget https://github.com/grpc/grpc-web/releases/download/$(PROTOC_GEN_GRPC_WEB_VERSION)/protoc-gen-grpc-web-$(PROTOC_GEN_GRPC_WEB_VERSION)-$(OS)-$(ARCH) \
			-O ./bin/protoc-gen-grpc-web \
		&& chmod +x ./bin/protoc-gen-grpc-web
	@GOPATH="$(shell pwd)/go" go install $(shell go list -f '{{range .Imports}}{{.}} {{end}}' tools.go)
	@npm i

.PHONY: generate
generate:
	go run -mod=mod ./cmd/generate
	cp -r ../proto-gen/web ../frontend/proto-gen-web
