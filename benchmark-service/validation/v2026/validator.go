package v2026

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	operatorAND = "and"
	operatorOR  = "or"
	sortAsc     = "asc"
	sortDesc    = "desc"
)

var (
	baseDate  = time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
	urlRegexp = regexp.MustCompile(`^http[s]?://farm\d\.static\.flickr\.com/\d+/.+\.jpg$`)
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

type Result struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Date string  `json:"date"`
	Url  string  `json:"url"`
}

type Response struct {
	Tag     string   `json:"tag"`
	Results []Result `json:"results"`
}

type Validator struct {
	logger        *slog.Logger
	geotagsByName map[string][]*Geotag
}

func NewValidator(logger *slog.Logger) *Validator {
	validator := &Validator{logger: logger}

	f, err := os.Open("data/v2023.json")
	if err != nil {
		logger.Error("[ERROR] cannot use validator.v2026", "err", err)
		return validator
	}
	defer f.Close()

	var tags []*Tag
	if err := json.NewDecoder(f).Decode(&tags); err != nil {
		logger.Error("failed to decode bytes to json", "err", err)
	}
	validator.geotagsByName = geotagsByName(tags)
	return validator
}

func NewValidatorWithTags(logger *slog.Logger, tags []*Tag) *Validator {
	return &Validator{
		logger:        logger,
		geotagsByName: geotagsByName(tags),
	}
}

func geotagsByName(tags []*Tag) map[string][]*Geotag {
	mp := make(map[string][]*Geotag, len(tags))
	for _, tag := range tags {
		mp[tag.TagName] = tag.Geotags
	}
	return mp
}

func (v *Validator) Validate(uri *url.URL, b []byte) error {
	var resp Response
	if err := json.Unmarshal(b, &resp); err != nil {
		v.logger.Info("failed to unmarshal", "err", err)
		return errors.New("json: invalid json format")
	}

	tags, operator, sortOrder, err := parseQuery(uri)
	if err != nil {
		return err
	}
	if resp.Tag != strings.Join(tags, ",") {
		return errors.New("tag: incorrect tag name")
	}
	if len(resp.Results) == 0 {
		return errors.New("Results: the length of Results must not be 0")
	}
	if len(resp.Results) > 100 {
		return errors.New("Results: the length of Results must be less than 100 or 100")
	}
	if err := validateResultShape(resp.Results, sortOrder); err != nil {
		return err
	}

	expected, err := v.expectedResults(tags, operator, sortOrder)
	if err != nil {
		return err
	}
	if len(expected) == 0 {
		return errors.New("Results: expected result must not be 0")
	}
	if len(expected) != len(resp.Results) {
		return fmt.Errorf("Results: the length of Results must be %d", len(expected))
	}
	for i := range resp.Results {
		if err := compareResult(i, expected[i], resp.Results[i]); err != nil {
			return err
		}
	}
	return nil
}

func parseQuery(uri *url.URL) ([]string, string, string, error) {
	q := uri.Query()
	tags := q["tag"]
	if len(tags) == 0 {
		return nil, "", "", errors.New("tag: tag must be set")
	}
	for _, tag := range tags {
		if tag == "" {
			return nil, "", "", errors.New("tag: tag must not be empty")
		}
	}

	operator := q.Get("tagOperator")
	if len(tags) == 1 {
		if operator == "" {
			operator = operatorOR
		}
	} else if operator == "" {
		return nil, "", "", errors.New("tagOperator: tagOperator must be set for multiple tags")
	}
	if operator != operatorAND && operator != operatorOR {
		return nil, "", "", errors.New("tagOperator: tagOperator must be and or or")
	}

	sortOrder := q.Get("sortOrder")
	if sortOrder == "" {
		sortOrder = sortDesc
	}
	if sortOrder != sortAsc && sortOrder != sortDesc {
		return nil, "", "", errors.New("sortOrder: sortOrder must be asc or desc")
	}
	return tags, operator, sortOrder, nil
}

func validateResultShape(results []Result, sortOrder string) error {
	for i, res := range results {
		if !urlRegexp.MatchString(res.Url) {
			return fmt.Errorf("Results[%d].Url: invalid format(expect %s)", i, urlRegexp.String())
		}
		if _, err := time.Parse("2006-01-02 15:04:05", res.Date); err != nil {
			return fmt.Errorf("Results[%d].Date: invalid format(expect yyyy-mm-dd hh:mm:dd)", i)
		}
	}
	for i := range results[:len(results)-1] {
		left, _ := time.Parse("2006-01-02 15:04:05", results[i].Date)
		right, _ := time.Parse("2006-01-02 15:04:05", results[i+1].Date)
		if sortOrder == sortDesc && left.Before(right) {
			return errors.New("Geotags: the order of Geotags must be desc by date")
		}
		if sortOrder == sortAsc && left.After(right) {
			return errors.New("Geotags: the order of Geotags must be asc by date")
		}
	}
	return nil
}

func (v *Validator) expectedResults(tags []string, operator, sortOrder string) ([]Result, error) {
	keyCounts := map[string]int{}
	geotagsByKey := map[string]*Geotag{}
	for _, tag := range tags {
		geotags, ok := v.geotagsByName[tag]
		if !ok {
			return nil, fmt.Errorf("tag: tag name is not found: %s", tag)
		}
		seenInTag := map[string]struct{}{}
		for _, geotag := range geotags {
			result := resultFromGeotag(geotag)
			key := resultKey(result)
			if _, ok := seenInTag[key]; ok {
				continue
			}
			seenInTag[key] = struct{}{}
			keyCounts[key]++
			geotagsByKey[key] = geotag
		}
	}

	expected := make([]Result, 0)
	for key, count := range keyCounts {
		if operator == operatorAND && count != len(tags) {
			continue
		}
		expected = append(expected, resultFromGeotag(geotagsByKey[key]))
	}
	sort.Slice(expected, func(i, j int) bool {
		left, _ := time.Parse("2006-01-02 15:04:05", expected[i].Date)
		right, _ := time.Parse("2006-01-02 15:04:05", expected[j].Date)
		if left.Equal(right) {
			return expected[i].Url < expected[j].Url
		}
		if sortOrder == sortAsc {
			return left.Before(right)
		}
		return left.After(right)
	})
	if len(expected) > 100 {
		expected = expected[:100]
	}
	return expected, nil
}

func resultFromGeotag(geotag *Geotag) Result {
	date := baseDate.Add(time.Duration(geotag.Elapsed) * time.Second)
	return Result{
		Lat:  geotag.Latitude,
		Lon:  geotag.Longitude,
		Date: date.Format("2006-01-02 15:04:05"),
		Url:  fmt.Sprintf("http://farm%d.static.flickr.com%s", geotag.FarmNum, geotag.Directory),
	}
}

func resultKey(result Result) string {
	return result.Url + result.Date
}

func compareResult(i int, expected, actual Result) error {
	if expected.Date != actual.Date {
		return fmt.Errorf("Results[%d].Date: incorrect date", i)
	}
	if !equalFloat(expected.Lat, actual.Lat) {
		return fmt.Errorf("Results[%d].Lat: incorrect latitude", i)
	}
	if !equalFloat(expected.Lon, actual.Lon) {
		return fmt.Errorf("Results[%d].Lon: incorrect longitude", i)
	}
	if expected.Url != actual.Url {
		return fmt.Errorf("Results[%d].Url: incorrect url", i)
	}
	return nil
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
