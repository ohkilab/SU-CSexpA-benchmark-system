module github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service

go 1.20

replace github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go => ../proto-gen/go

require (
	github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.54.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
