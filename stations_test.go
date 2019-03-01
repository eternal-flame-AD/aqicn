package aqicn

import (
	"testing"
)

func TestStationSearch(t *testing.T) {
	client := Client{Token: "demo"}
	res, err := client.SearchStation(39.379436, 116.091230, 40.235643, 116.784382)
	if err != nil {
		t.Fatal(err)
	}
	_ = res
	// fmt.Println(res)
}
