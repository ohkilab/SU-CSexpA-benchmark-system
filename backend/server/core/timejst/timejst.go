package timejst

import "time"

var jst *time.Location

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

func Now() time.Time {
	return time.Now().In(jst)
}
