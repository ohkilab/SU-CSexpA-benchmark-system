package validation

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/url"
	"time"
)

// https://ohkilab.github.io/SU-CSexpA/content/part3/part3_final_assignment/final_assignment_details.html#geotag-csv
type Response2022 struct {
	Tag     string `json:"tag"`
	Geotags []struct {
		Lat  float64 `json:"lat"`
		Lon  float64 `json:"lon"`
		Date string  `json:"date"`
		Url  string  `json:"url"`
	} `json:"geotags"`
}

func Validate2022(uri *url.URL, r io.ReadCloser) error {
	defer func() {
		_, _ = io.Copy(io.Discard, r)
		r.Close()
	}()
	var resp Response2022
	if err := json.NewDecoder(r).Decode(&resp); err != nil {
		log.Println(err)
		return errors.New("json: invalid json format")
	}
	tag := uri.Query().Get("tag")
	if resp.Tag != tag {
		return errors.New("tag: incorrect tag name")
	}
	if len(resp.Geotags) == 0 {
		return errors.New("Geotags: the length of Geotags must not be 0")
	}
	if len(resp.Geotags) > 100 {
		return errors.New("Geotags: the length of Geotags must be less than 100 or 100")
	}
	for i := range resp.Geotags[:len(resp.Geotags)-1] {
		left, _ := time.Parse("2006-01-02UTC", resp.Geotags[i].Date)
		right, _ := time.Parse("2006-01-02UTC", resp.Geotags[i+1].Date)
		// left, right := resp.Geotags[i].Date, resp.Geotags[i+1].Date
		if left.After(right) {
			return errors.New("Geotags: the order of Geotags must be desc by asc")
		}
	}
	for _, res := range resp.Geotags {
		if _, err := url.ParseRequestURI(res.Url); err != nil {
			return errors.New("url: invalid url")
		}
	}

	return nil
}
