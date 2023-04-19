module github.com/ohkilab/SU-CSexpA-benchmark-system/backend

go 1.20

replace github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go => ../proto-gen/go

require (
	entgo.io/ent v0.12.1
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.54.0
)

require (
	ariga.io/atlas v0.10.0 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
