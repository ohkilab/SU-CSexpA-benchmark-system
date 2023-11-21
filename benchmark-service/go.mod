module github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service

go 1.21

replace github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go => ../proto-gen/go

require (
	github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go v0.0.0-20231121052944-e7c394a699cc
	github.com/stretchr/testify v1.8.2
	golang.org/x/sync v0.5.0
	google.golang.org/grpc v1.59.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
