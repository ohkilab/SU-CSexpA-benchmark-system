package validation

import (
	"net/url"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ValidateFunc func(uri *url.URL, b []byte) error

// NOTE: これ非常に良くない
// Year の中で予選と本選があり、その validator の spec が同じことが保証される必要がある
// かといってコンテストごとに validator を定義するのはだるいので、とりあえずこの仕様に任せることにする・・
func Detect(year int) (ValidateFunc, error) {
	switch year {
	case 2022:
		return validate2022, nil
	case 2023:
		return validate2023, nil
	default:
		return nil, status.Errorf(codes.InvalidArgument, "the validator of year %d is not found", year)
	}
}
