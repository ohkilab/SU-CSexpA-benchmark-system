.PHONY: tools/install
tools/install:
	@npm i
	@GOPATH="$(shell pwd)/go" go install $(shell go list -f '{{range .Imports}}{{.}} {{end}}' tools.go)

.PHONY: generate
generate:
	go run -mod=mod ./cmd/generate