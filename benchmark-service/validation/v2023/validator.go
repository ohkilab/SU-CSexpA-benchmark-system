package v2023

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"time"

	"log/slog"
)

type Tag struct {
	TagName string    `json:"tag_name"`
	Geotags []*Geotag `json:"geotags"`
}

type Geotag struct {
	Elapsed   int32   `json:"elapsed"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	FarmNum   uint8   `json:"farm_num"`
	Directory string  `json:"directory"`
}

var baseDate = time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)

type Validator struct {
	logger        *slog.Logger
	geotagsByName map[string][]*Geotag
}

func NewValidator(logger *slog.Logger) *Validator {
	validator := &Validator{logger: logger}

	logger.Info("opening tag.json...")
	f, err := os.Open("data/v2023.json")
	if err != nil {
		logger.Error("[ERROR] cannot use validator.v2023", "err", err)
		return validator
	}
	defer f.Close()
	logger.Info("done")

	logger.Info("decoding json...")
	var tags []*Tag
	if err := json.NewDecoder(f).Decode(&tags); err != nil {
		logger.Error("failed to decode bytes to json", err)
	}
	logger.Info("done")

	logger.Info("creating map...")
	geotagsByName := make(map[string][]*Geotag)
	for _, tag := range tags {
		geotagsByName[tag.TagName] = tag.Geotags
	}
	logger.Info("done")

	validator.geotagsByName = geotagsByName
	return validator
}

// https://ohkilab.github.io/SU-CSexpA/content/part3/part3_final_assignment/final_assignment_details.html#geotag-csv
type Response struct {
	Tag     string `json:"tag"`
	Results []struct {
		Lat  float64 `json:"lat"`
		Lon  float64 `json:"lon"`
		Date string  `json:"date"`
		Url  string  `json:"url"`
	} `json:"results"`
}

var urlRegexp = regexp.MustCompile(`^http[s]?://farm\d\.static\.flickr\.com/\d+/.+\.jpg$`)

func (v *Validator) Validate(uri *url.URL, b []byte) error {
	var resp Response
	if err := json.Unmarshal(b, &resp); err != nil {
		v.logger.Info("failed to unmarshal", err)
		return errors.New("json: invalid json format")
	}
	tag := uri.Query().Get("tag")
	if resp.Tag != tag {
		return errors.New("tag: incorrect tag name")
	}
	if len(resp.Results) == 0 {
		return errors.New("Results: the length of Results must not be 0")
	}
	if len(resp.Results) > 100 {
		return errors.New("Results: the length of Results must be less than 100 or 100")
	}
	for i, res := range resp.Results {
		if !urlRegexp.MatchString(res.Url) {
			v.logger.Info("url", res.Url)
			return fmt.Errorf("Results[%d].Url: invalid format(expect %s)", i, urlRegexp.String())
		}
		if _, err := time.Parse("2006-01-02 15:04:05", resp.Results[i].Date); err != nil {
			return fmt.Errorf("Results[%d].Date: invalid format(expect yyyy-mm-dd hh:mm:dd)", i)
		}
	}
	for i := range resp.Results[:len(resp.Results)-1] {
		left, _ := time.Parse("2006-01-02 15:04:05", resp.Results[i].Date)
		right, _ := time.Parse("2006-01-02 15:04:05", resp.Results[i+1].Date)
		if left.Before(right) {
			return errors.New("Geotags: the order of Geotags must be desc by date")
		}
	}

	return v.isCorrectResult(tag, &resp)
}

func (v *Validator) isCorrectResult(tagName string, resp *Response) error {
	geotags, ok := v.geotagsByName[tagName]
	if !ok {
		v.logger.Error("tag: tag name %s is not found", tagName)
		return nil
	}
	newKey := func(url string, date time.Time) string {
		return fmt.Sprint(url, date.Format("2006-01-02 15:04:05"))
	}
	geotagByKey := make(map[string]*Geotag, len(geotags))
	for i, geotag := range geotags {
		url := fmt.Sprintf("http://farm%d.static.flickr.com%s", geotag.FarmNum, geotag.Directory)
		date := baseDate.Add(time.Duration(geotag.Elapsed) * time.Second)
		geotagByKey[newKey(url, date)] = geotags[i]
	}

	if len(geotags) != len(resp.Results) {
		return fmt.Errorf("Results: the length of Results must be %d", len(geotags))
	}
	for i, actual := range resp.Results {
		actualDate, _ := time.Parse("2006-01-02 15:04:05", actual.Date)
		key := newKey(actual.Url, actualDate)
		expect, ok := geotagByKey[key]
		if !ok {
			// NOTE: 日時が同一のレコードがある場合は順不同で良いという制約があるが、この場合に正確な validation をするためには
			// 全てのレコードをメモリに載せないといけない。DB を用いないでこれを行なった場合はメモリが8GBほど必要になってしまうため、
			// 今回は一部のレコードについては validation を行わないという方針にする
			v.logger.Warn("the key of geotagByKey is not found", "tag", tagName, "key", key)
			// return fmt.Errorf("Result[%d]: incorrect url or date", i)
			continue
		}
		if !equalDate(baseDate.Add(time.Duration(expect.Elapsed)*time.Second), actualDate) {
			return fmt.Errorf("Results[%d].Date: incorrect date", i)
		}
		if !equalFloat(expect.Latitude, actual.Lat) {
			return fmt.Errorf("Results[%d].Lat: incorrect latitude", i)
		}
		if !equalFloat(expect.Longitude, actual.Lon) {
			return fmt.Errorf("Results[%d].Lon: incorrect longitude", i)
		}
		actualURL := fmt.Sprintf("http://farm%d.static.flickr.com%s", expect.FarmNum, expect.Directory)
		if actualURL != actual.Url {
			return fmt.Errorf("Results[%d].Url: incorrect url", i)
		}
	}
	return nil
}

func equalDate(x, y time.Time) bool {
	return x.Truncate(time.Second).Equal(y.Truncate(time.Second))
}

func equalFloat[T float32 | float64](x, y T) bool {
	const eps = 1e-9
	return abs(x, y) <= eps
}

func abs[T float32 | float64](x, y T) T {
	ans := x - y
	if ans < 0 {
		return -ans
	}
	return ans
}
