package aqicn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

// Station represents a SearchStation response
type Station struct {
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	StationID int     `json:"uid"`
	AQI       string  `json:"aqi"`
}

// SearchStation searches available stations in a given geological range
func (c Client) SearchStation(lat1, lng1, lat2, lng2 float64) ([]Station, error) {
	token := url.QueryEscape(c.Token)
	url := fmt.Sprintf("https://api.waqi.info/map/bounds/?latlng=%.6f,%.6f,%.6f,%.6f&token=%s", lat1, lng1, lat2, lng2, token)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data, err := getData(r)
	if err != nil {
		return nil, err
	}
	res := make([]Station, 0)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}
