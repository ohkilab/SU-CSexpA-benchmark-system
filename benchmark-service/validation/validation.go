package validation

import (
	"net/url"

	v2023 "github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation/v2023"
	v2026 "github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation/v2026"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"log/slog"
)

type Validator interface {
	Validate(uri *url.URL, b []byte) error
}

// NOTE: これ非常に良くない
// 将来的には validation は single の go にして、github repository から download するなどしたい
func NewValidator(logger *slog.Logger) map[pb.Validator]Validator {
	v2023Validator := v2023.NewValidator(logger)
	v2026Validator := v2026.NewValidator(logger)

	return map[pb.Validator]Validator{
		pb.Validator_V2023: v2023Validator,
		pb.Validator(2):    v2026Validator,
	}
}
