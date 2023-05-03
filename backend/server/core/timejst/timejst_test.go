package timejst

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Now(t *testing.T) {
	nowUtc := time.Now().UTC()
	nowJst := Now()
	assert.Equal(t, nowJst.Location().String(), "Asia/Tokyo")
	assert.Equal(t, nowUtc.Add(9*time.Hour).Format("2006-01-02 15:04:05"), nowJst.Format("2006-01-02 15:04:05"))
}
