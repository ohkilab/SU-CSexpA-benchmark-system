package validation

import (
	"net/url"

	v2023 "github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation/v2023"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"log/slog"
)

type Validator interface {
	Validate(uri *url.URL, b []byte) error
}

// NOTE: これ非常に良くない
// 将来的には validation は single の go にして、github repository から download するなどしたい
func NewValidator(logger *slog.Logger) map[string]Validator {
	v2023Validator := v2023.NewValidator(logger)

	return map[string]Validator{
		pb.Validator_V2023.String(): v2023Validator,
	}
}
