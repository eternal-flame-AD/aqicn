package aqicn

import (
	"testing"
)

func TestCityFeed(t *testing.T) {
	client := Client{"demo"}
	res, err := client.CityFeed("shanghai")
	if err != nil {
		t.Fatal(err)
	}
	_ = res
	//fmt.Println(res)
}
