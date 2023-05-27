package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"time"
)

// https://ohkilab.github.io/SU-CSexpA/content/part3/part3_final_assignment/final_assignment_details.html#geotag-csv
type Response2023 struct {
	Tag     string `json:"tag"`
	Results []struct {
		Lat  float64 `json:"lat"`
		Lon  float64 `json:"lon"`
		Date string  `json:"date"`
		Url  string  `json:"url"`
	} `json:"results"`
}

var urlRegexp = regexp.MustCompile(`^http[s]?://farm\d\.static\.flickr\.com/\d+/.+\.jpg$`)

func validate2023(uri *url.URL, b []byte) error {
	var resp Response2023
	if err := json.Unmarshal(b, &resp); err != nil {
		log.Println(err)
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
	for i := range resp.Results[:len(resp.Results)-1] {
		left, _ := time.Parse("2006-01-02 15:04:05", resp.Results[i].Date)
		right, _ := time.Parse("2006-01-02 15:04:05", resp.Results[i+1].Date)
		// left, right := resp.Geotags[i].Date, resp.Geotags[i+1].Date
		if left.After(right) {
			return errors.New("Geotags: the order of Geotags must be desc by date")
		}
	}
	for i, res := range resp.Results {
		if !urlRegexp.MatchString(res.Url) {
			log.Println(res.Url)
			return fmt.Errorf("Results[%d].Url: invalid format", i)
		}
	}

	return nil
}
