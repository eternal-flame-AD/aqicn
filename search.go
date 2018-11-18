package aqicn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SearchCity searches stations of a given city name
func (c Client) SearchCity(keyword string) ([]SearchFeed, error) {
	keyword = url.PathEscape(keyword)
	token := url.QueryEscape(c.Token)
	resp, err := http.Get(fmt.Sprintf("https://api.waqi.info/search/?keyword=%s&token=%s", keyword, token))
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
	res := make([]SearchFeed, 0)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}
