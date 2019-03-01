package aqicn

import "net/http"

// Client is AQICN Client
type Client struct {
	Token  string
	Client *http.Client
}
