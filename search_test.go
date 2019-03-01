package aqicn

import (
	"testing"
)

func TestCitySearch(t *testing.T) {
	client := Client{Token: "demo"}
	res, err := client.SearchCity("s")
	if err != nil {
		t.Fatal(err)
	}
	_ = res
}
