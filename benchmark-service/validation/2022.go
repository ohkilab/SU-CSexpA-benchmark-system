package validation

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"time"
)

// https://ohkilab.github.io/SU-CSexpA/content/part3/part3_final_assignment/final_assignment_details.html#geotag-csv
type Response2022 struct {
	Tag     string `json:"tag"`
	Results []struct {
		Lat  float64   `json:"lat"`
		Lon  float64   `json:"lon"`
		Date time.Time `json:"date"`
		Url  string    `json:"url"`
	} `json:"results"`
}

func Validate2022(uri *url.URL, r io.ReadCloser) error {
	defer r.Close()
	var resp *Response2022
	if err := json.NewDecoder(r).Decode(resp); err != nil {
		return errors.New("json: invalid json format")
	}
	tag := uri.Query().Get("tag")
	if resp.Tag != tag {
		return errors.New("tag: incorrect tag name")
	}
	if len(resp.Results) == 0 {
		return errors.New("results: the length of results must not be 0")
	}
	if len(resp.Results) > 100 {
		return errors.New("results: the length of results must be less than 100 or 100")
	}
	for i := range resp.Results[:len(resp.Results)-1] {
		if resp.Results[i].Date.Before(resp.Results[i+1].Date) {
			return errors.New("results: the order of results must be desc by date")
		}
	}
	for _, res := range resp.Results {
		if _, err := url.ParseRequestURI(res.Url); err != nil {
			return errors.New("url: invalid url")
		}
	}

	return nil
}
