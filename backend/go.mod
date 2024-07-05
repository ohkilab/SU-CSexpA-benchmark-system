module github.com/ohkilab/SU-CSexpA-benchmark-system/backend

go 1.22.3

replace github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go => ../proto-gen/go

replace github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service => ../benchmark-service

require (
	entgo.io/ent v0.13.1
	github.com/go-sql-driver/mysql v1.8.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service v0.0.0-00010101000000-000000000000
	github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go v0.0.0-20240214075016-ed326009ee8b
	github.com/pkg/errors v0.9.1
	github.com/samber/lo v1.39.0
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/crypto v0.24.0
	golang.org/x/exp v0.0.0-20240604190554-fc45aab8b7f8
	golang.org/x/net v0.26.0
	golang.org/x/sync v0.7.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
)

require (
	ariga.io/atlas v0.23.0 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-openapi/inflect v0.21.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.20.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
