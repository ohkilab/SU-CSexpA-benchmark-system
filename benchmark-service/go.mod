module github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service

go 1.21

replace github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go => ../proto-gen/go

require (
	github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go v0.0.0-20231206071327-6ea8f5f13f1c
	github.com/stretchr/testify v1.8.2
	golang.org/x/sync v0.6.0
	google.golang.org/grpc v1.61.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240213162025-012b6fc9bca9 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
