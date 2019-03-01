package aqicn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

// Here fetches the feed of the station nearest your location based on your IP address
func (c Client) Here() (*DetailFeed, error) {
	token := url.QueryEscape(c.Token)
	resp, err := c.get(fmt.Sprintf("https://api.waqi.info/feed/here/?token=%s", token))
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

// FeedAt fetches the feed of the station located nearest the given latitude and longitude
func (c Client) FeedAt(lat, lon float64) (*DetailFeed, error) {
	token := url.QueryEscape(c.Token)
	resp, err := c.get(fmt.Sprintf("https://api.waqi.info/feed/geo:%.6f;%.6f/?token=%s", lat, lon, token))
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
