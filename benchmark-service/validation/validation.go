package validation

import (
	"net/url"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ValidateFunc func(uri *url.URL, b []byte) error

// NOTE: これ非常に良くない
// 将来的には validation は single の go にして、github repository から download するなどしたい
func Detect(slug string) (ValidateFunc, error) {
	switch slug {
	case "test-contest":
		return validate2023, nil
	case "2023-qual":
		return validate2023, nil
	case "2023-final":
		return validate2023, nil
	default:
		return nil, status.Errorf(codes.InvalidArgument, "the validator the contest_slug is %s is not found", slug)
	}
}
