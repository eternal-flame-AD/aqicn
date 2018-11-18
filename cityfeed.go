package aqicn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// CityFeed fetches the feed of a given city name
func (c Client) CityFeed(city string) (*DetailFeed, error) {
	city = url.PathEscape(city)
	token := url.QueryEscape(c.Token)
	resp, err := http.Get(fmt.Sprintf("https://api.waqi.info/feed/%s/?token=%s", city, token))
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
	res := new(DetailFeed)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// StationFeed fetches the feed of a given station ID
func (c Client) StationFeed(id int) (*DetailFeed, error) {
	return c.CityFeed(fmt.Sprintf("@%d", id))
}
