package aqicn

import "net/http"

func (c Client) get(url string) (resp *http.Response, err error) {
	if c.Client == nil {
		c.Client = new(http.Client)
	}
	return c.Client.Get(url)
}
