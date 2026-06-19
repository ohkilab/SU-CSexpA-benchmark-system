package submit

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildBenchmarkURLSingleTag(t *testing.T) {
	actual := buildBenchmarkURL("http://example.com/search", "cat")

	assert.Equal(t, "http://example.com/search?tag=cat", actual)
}

func TestBuildBenchmarkURLQueryFragment(t *testing.T) {
	actual := buildBenchmarkURL("http://example.com/search", "tag=cat&tag=dog&tagOperator=and&sortOrder=desc")
	uri, err := url.Parse(actual)
	require.NoError(t, err)

	query := uri.Query()
	assert.Equal(t, "http://example.com/search", uri.Scheme+"://"+uri.Host+uri.Path)
	assert.Equal(t, []string{"cat", "dog"}, query["tag"])
	assert.Equal(t, "and", query.Get("tagOperator"))
	assert.Equal(t, "desc", query.Get("sortOrder"))
}

func TestBuildBenchmarkURLKeepsExistingQuery(t *testing.T) {
	actual := buildBenchmarkURL("http://example.com/search?token=abc", "tag=cat&tag=dog&tagOperator=or&sortOrder=asc")
	uri, err := url.Parse(actual)
	require.NoError(t, err)

	query := uri.Query()
	assert.Equal(t, "abc", query.Get("token"))
	assert.Equal(t, []string{"cat", "dog"}, query["tag"])
	assert.Equal(t, "or", query.Get("tagOperator"))
	assert.Equal(t, "asc", query.Get("sortOrder"))
}
