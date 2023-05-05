package validation

import (
	"bytes"
	"io"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testdata = `
{
	"tag": "陸上自衛隊",
	"geotags": [
	  {
		"lat": 35.285423,
		"lon": 138.84761,
		"date": "2012-08-21UTC",
		"url": "https://farm9.static.flickr.com/8282/7831339680_265efb968e.jpg"
	  },
	  {
		"lat": 35.285423,
		"lon": 138.84761,
		"date": "2012-08-21UTC",
		"url": "https://farm8.static.flickr.com/7124/7831342090_9383001e63.jpg"
	  }
	]
}`

func Test_validate2022(t *testing.T) {
	r := io.NopCloser(bytes.NewReader([]byte(testdata)))
	uri, _ := url.ParseRequestURI("http://localhost:8080/program?tag=陸上自衛隊")
	err := Validate2022(uri, r)
	assert.NoError(t, err)
}
