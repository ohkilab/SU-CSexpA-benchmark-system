package v2026

import (
	"encoding/json"
	"log/slog"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRegisteredSingleTagCase(t *testing.T) {
	rawURL := "http://localhost/search?tag=cat"
	results := []Result{
		testResult(30, 1),
		testResult(20, 2),
		testResult(10, 3),
	}
	validator := validatorWithCase(t, rawURL, results)
	body := mustResponse(t, "cat", results)

	assert.NoError(t, validator.Validate(mustURL(t, rawURL), body))
}

func TestValidateRegisteredMultipleTagsOR(t *testing.T) {
	rawURL := "http://localhost/search?tag=cat&tag=dog&tagOperator=or&sortOrder=desc"
	results := []Result{
		testResult(40, 4),
		testResult(30, 1),
		testResult(20, 2),
		testResult(10, 3),
	}
	validator := validatorWithCase(t, rawURL, results)
	body := mustResponse(t, "cat,dog", results)

	assert.NoError(t, validator.Validate(mustURL(t, rawURL), body))
}

func TestValidateRegisteredMultipleTagsAND(t *testing.T) {
	rawURL := "http://localhost/search?tag=cat&tag=dog&tagOperator=and&sortOrder=desc"
	results := []Result{testResult(20, 2)}
	validator := validatorWithCase(t, rawURL, results)
	body := mustResponse(t, "cat,dog", results)

	assert.NoError(t, validator.Validate(mustURL(t, rawURL), body))
}

func TestValidateSortOrderAsc(t *testing.T) {
	rawURL := "http://localhost/search?tag=cat&sortOrder=asc"
	results := []Result{
		testResult(10, 3),
		testResult(20, 2),
		testResult(30, 1),
	}
	validator := validatorWithCase(t, rawURL, results)
	body := mustResponse(t, "cat", results)

	assert.NoError(t, validator.Validate(mustURL(t, rawURL), body))
}

func TestValidateRejectsUnregisteredQuery(t *testing.T) {
	registeredURL := "http://localhost/search?tag=cat"
	requestURL := "http://localhost/search?tag=dog"
	results := []Result{testResult(30, 1)}
	validator := validatorWithCase(t, registeredURL, results)
	body := mustResponse(t, "dog", results)

	err := validator.Validate(mustURL(t, requestURL), body)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "query is not found")
}

func TestValidateRejectsIncorrectResult(t *testing.T) {
	rawURL := "http://localhost/search?tag=cat"
	results := []Result{
		testResult(30, 1),
		testResult(20, 2),
		testResult(10, 3),
	}
	validator := validatorWithCase(t, rawURL, results)
	actual := append([]Result(nil), results...)
	actual[1].Lat = 99
	body := mustResponse(t, "cat", actual)

	err := validator.Validate(mustURL(t, rawURL), body)
	if err == nil {
		t.Fatal("expected error")
	}
	assert.Contains(t, err.Error(), "incorrect latitude")
}

func TestValidateRejectsInvalidQuery(t *testing.T) {
	validator := validatorWithCase(t, "http://localhost/search?tag=cat", []Result{testResult(30, 1)})
	tests := []string{
		"http://localhost/search?tag=cat&tag=dog",
		"http://localhost/search?tag=cat&tag=dog&tagOperator=x",
		"http://localhost/search?tag=cat&sortOrder=x",
	}
	for _, rawURL := range tests {
		t.Run(rawURL, func(t *testing.T) {
			body := mustResponse(t, "cat", []Result{testResult(30, 1)})
			assert.Error(t, validator.Validate(mustURL(t, rawURL), body))
		})
	}
}

func TestValidateRejectsWrongOrder(t *testing.T) {
	rawURL := "http://localhost/search?tag=cat"
	expected := []Result{
		testResult(30, 1),
		testResult(20, 2),
		testResult(10, 3),
	}
	actual := []Result{
		testResult(10, 3),
		testResult(20, 2),
		testResult(30, 1),
	}
	validator := validatorWithCase(t, rawURL, expected)
	body := mustResponse(t, "cat", actual)

	err := validator.Validate(mustURL(t, rawURL), body)
	if err == nil {
		t.Fatal("expected error")
	}
	assert.Contains(t, err.Error(), "desc")
}

func TestCanonicalQueryMatchesGeneratedCaseShape(t *testing.T) {
	uri := mustURL(t, "http://localhost/search?tag=cat&tag=dog&tagOperator=and&sortOrder=asc")
	tags, operator, sortOrder, err := parseQuery(uri)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "sortOrder=asc&tag=cat&tag=dog&tagOperator=and", canonicalQuery(tags, operator, sortOrder))
}

func validatorWithCase(t *testing.T, rawURL string, results []Result) *Validator {
	t.Helper()
	uri := mustURL(t, rawURL)
	tags, operator, sortOrder, err := parseQuery(uri)
	if err != nil {
		t.Fatal(err)
	}
	return NewValidatorWithCases(slog.Default(), []*Case{
		{
			Query:       canonicalQuery(tags, operator, sortOrder),
			Tags:        tags,
			TagOperator: operator,
			SortOrder:   sortOrder,
			Results:     results,
		},
	})
}

func testResult(elapsed int32, id int) Result {
	return Result{
		Lat:  float64(id),
		Lon:  float64(id) + 0.1,
		Date: dateFromElapsed(elapsed),
		Url:  "http://farm1.static.flickr.com/100/photo" + string(rune('a'+id)) + ".jpg",
	}
}

func dateFromElapsed(elapsed int32) string {
	switch elapsed {
	case 10:
		return "2012-01-01 00:00:10"
	case 20:
		return "2012-01-01 00:00:20"
	case 30:
		return "2012-01-01 00:00:30"
	case 40:
		return "2012-01-01 00:00:40"
	default:
		return "2012-01-01 00:00:00"
	}
}

func mustURL(t *testing.T, rawURL string) *url.URL {
	t.Helper()
	uri, err := url.ParseRequestURI(rawURL)
	if err != nil {
		t.Fatal(err)
	}
	return uri
}

func mustResponse(t *testing.T, tag string, results []Result) []byte {
	t.Helper()
	b, err := json.Marshal(Response{Tag: tag, Results: results})
	if err != nil {
		t.Fatal(err)
	}
	return b
}
