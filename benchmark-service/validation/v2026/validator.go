package v2026

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	operatorAND = "and"
	operatorOR  = "or"
	sortAsc     = "asc"
	sortDesc    = "desc"
)

var urlRegexp = regexp.MustCompile(`^http[s]?://farm\d\.static\.flickr\.com/\d+/.+\.jpg$`)

type V2026Data struct {
	Version string  `json:"version"`
	Cases   []*Case `json:"cases"`
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

type Case struct {
	Query       string   `json:"query"`
	Tags        []string `json:"tags"`
	TagOperator string   `json:"tag_operator"`
	SortOrder   string   `json:"sort_order"`
	Results     []Result `json:"results"`
}

type Validator struct {
	logger       *slog.Logger
	casesByQuery map[string]*Case
}

func NewValidator(logger *slog.Logger) *Validator {
	validator := &Validator{logger: logger}

	f, err := os.Open("data/v2026.json")
	if err != nil {
		logger.Error("[ERROR] cannot use validator.v2026", "err", err)
		return validator
	}
	defer f.Close()

	var data V2026Data
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		logger.Error("failed to decode bytes to json", "err", err)
	}
	validator.casesByQuery = casesByQuery(data.Cases)
	return validator
}

func NewValidatorWithCases(logger *slog.Logger, cases []*Case) *Validator {
	return &Validator{
		logger:       logger,
		casesByQuery: casesByQuery(cases),
	}
}

func casesByQuery(cases []*Case) map[string]*Case {
	mp := make(map[string]*Case, len(cases))
	for _, c := range cases {
		mp[c.Query] = c
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

	query := canonicalQuery(tags, operator, sortOrder)
	c, ok := v.casesByQuery[query]
	if !ok {
		return fmt.Errorf("case: query is not found: %s", query)
	}
	expected := c.Results
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

func canonicalQuery(tags []string, operator, sortOrder string) string {
	values := url.Values{}
	values.Set("sortOrder", sortOrder)
	for _, tag := range tags {
		values.Add("tag", tag)
	}
	if len(tags) > 1 {
		values.Set("tagOperator", operator)
	}
	return values.Encode()
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
		if compareResultOrder(results[i], results[i+1], sortOrder) > 0 {
			return fmt.Errorf("Geotags: the order of Geotags must be %s by date and asc by url", sortOrder)
		}
	}
	return nil
}

func compareResultOrder(left, right Result, sortOrder string) int {
	leftDate, _ := time.Parse("2006-01-02 15:04:05", left.Date)
	rightDate, _ := time.Parse("2006-01-02 15:04:05", right.Date)
	if !leftDate.Equal(rightDate) {
		if sortOrder == sortAsc {
			if leftDate.Before(rightDate) {
				return -1
			}
			return 1
		}
		if leftDate.After(rightDate) {
			return -1
		}
		return 1
	}
	if left.Url < right.Url {
		return -1
	}
	if left.Url > right.Url {
		return 1
	}
	return 0
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
