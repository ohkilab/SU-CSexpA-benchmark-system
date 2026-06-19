package v2026

import (
	"encoding/json"
	"log/slog"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateSingleTagCompatible(t *testing.T) {
	validator := NewValidatorWithTags(slog.Default(), testTags())
	uri := mustURL(t, "http://localhost/search?tag=cat")
	body := mustResponse(t, "cat", []Result{
		resultFromGeotag(testGeotag(30, 1)),
		resultFromGeotag(testGeotag(20, 2)),
		resultFromGeotag(testGeotag(10, 3)),
	})

	assert.NoError(t, validator.Validate(uri, body))
}

func TestValidateMultipleTagsOR(t *testing.T) {
	validator := NewValidatorWithTags(slog.Default(), testTags())
	uri := mustURL(t, "http://localhost/search?tag=cat&tag=dog&tagOperator=or&sortOrder=desc")
	body := mustResponse(t, "cat,dog", []Result{
		resultFromGeotag(testGeotag(40, 4)),
		resultFromGeotag(testGeotag(30, 1)),
		resultFromGeotag(testGeotag(20, 2)),
		resultFromGeotag(testGeotag(10, 3)),
	})

	assert.NoError(t, validator.Validate(uri, body))
}

func TestValidateMultipleTagsAND(t *testing.T) {
	validator := NewValidatorWithTags(slog.Default(), testTags())
	uri := mustURL(t, "http://localhost/search?tag=cat&tag=dog&tagOperator=and&sortOrder=desc")
	body := mustResponse(t, "cat,dog", []Result{
		resultFromGeotag(testGeotag(20, 2)),
	})

	assert.NoError(t, validator.Validate(uri, body))
}

func TestValidateSortOrderAsc(t *testing.T) {
	validator := NewValidatorWithTags(slog.Default(), testTags())
	uri := mustURL(t, "http://localhost/search?tag=cat&sortOrder=asc")
	body := mustResponse(t, "cat", []Result{
		resultFromGeotag(testGeotag(10, 3)),
		resultFromGeotag(testGeotag(20, 2)),
		resultFromGeotag(testGeotag(30, 1)),
	})

	assert.NoError(t, validator.Validate(uri, body))
}

func TestValidateLimitsExpectedResultsToFirst100(t *testing.T) {
	geotags := make([]*Geotag, 0, 101)
	results := make([]Result, 0, 100)
	for i := 101; i >= 1; i-- {
		geotags = append(geotags, testGeotag(int32(i), i))
		if len(results) < 100 {
			results = append(results, resultFromGeotag(testGeotag(int32(i), i)))
		}
	}
	validator := NewValidatorWithTags(slog.Default(), []*Tag{{TagName: "cat", Geotags: geotags}})
	uri := mustURL(t, "http://localhost/search?tag=cat")
	body := mustResponse(t, "cat", results)

	assert.NoError(t, validator.Validate(uri, body))
}

func TestValidateRejectsIncorrectResult(t *testing.T) {
	validator := NewValidatorWithTags(slog.Default(), testTags())
	uri := mustURL(t, "http://localhost/search?tag=cat")
	results := []Result{
		resultFromGeotag(testGeotag(30, 1)),
		resultFromGeotag(testGeotag(20, 2)),
		resultFromGeotag(testGeotag(10, 3)),
	}
	results[1].Lat = 99
	body := mustResponse(t, "cat", results)

	err := validator.Validate(uri, body)
	if err == nil {
		t.Fatal("expected error")
	}
	assert.Contains(t, err.Error(), "incorrect latitude")
}

func TestValidateRejectsInvalidQuery(t *testing.T) {
	validator := NewValidatorWithTags(slog.Default(), testTags())
	tests := []string{
		"http://localhost/search?tag=cat&tag=dog",
		"http://localhost/search?tag=cat&tag=dog&tagOperator=x",
		"http://localhost/search?tag=cat&sortOrder=x",
		"http://localhost/search?tag=unknown",
	}
	for _, rawURL := range tests {
		t.Run(rawURL, func(t *testing.T) {
			body := mustResponse(t, "cat", []Result{resultFromGeotag(testGeotag(30, 1))})
			assert.Error(t, validator.Validate(mustURL(t, rawURL), body))
		})
	}
}

func TestValidateRejectsWrongOrder(t *testing.T) {
	validator := NewValidatorWithTags(slog.Default(), testTags())
	uri := mustURL(t, "http://localhost/search?tag=cat")
	body := mustResponse(t, "cat", []Result{
		resultFromGeotag(testGeotag(10, 3)),
		resultFromGeotag(testGeotag(20, 2)),
		resultFromGeotag(testGeotag(30, 1)),
	})

	err := validator.Validate(uri, body)
	if err == nil {
		t.Fatal("expected error")
	}
	assert.Contains(t, err.Error(), "desc")
}

func testTags() []*Tag {
	return []*Tag{
		{
			TagName: "cat",
			Geotags: []*Geotag{
				testGeotag(30, 1),
				testGeotag(20, 2),
				testGeotag(10, 3),
			},
		},
		{
			TagName: "dog",
			Geotags: []*Geotag{
				testGeotag(40, 4),
				testGeotag(20, 2),
			},
		},
	}
}

func testGeotag(elapsed int32, id int) *Geotag {
	return &Geotag{
		Elapsed:   elapsed,
		Latitude:  float64(id),
		Longitude: float64(id) + 0.1,
		FarmNum:   1,
		Directory: "/100/photo" + string(rune('a'+id)) + ".jpg",
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
