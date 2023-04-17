.PHONY: tools/install
tools/install:
	npm install -g protoc-gen-ts
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: generate
generate:
	go run -mod=mod ./cmd/generate