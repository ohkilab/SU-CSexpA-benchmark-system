module github.com/ohkilab/SU-CSexpA-benchmark-system/backend

go 1.20

replace github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go => ../proto-gen/go

replace github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service => ../benchmark-service

require (
	entgo.io/ent v0.12.3
	github.com/go-sql-driver/mysql v1.7.1
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service v0.0.0-00010101000000-000000000000
	github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/samber/lo v1.38.1
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.2
	golang.org/x/crypto v0.11.0
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1
	golang.org/x/net v0.12.0
	golang.org/x/sync v0.3.0
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	ariga.io/atlas v0.12.0 // indirect
	cloud.google.com/go/compute v1.19.3 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.17.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zclconf/go-cty v1.13.2 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230717213848-3f92550aa753 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
