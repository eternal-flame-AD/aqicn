package aqicn

import (
	"encoding/json"
	"testing"
	"time"
)

var mockObsTime = []byte(`{"time":{
	"v":1481396400,
	"s":"2016-12-10 19:00:00",
	"tz":"-06:00"
}}`)

func TestObsTimeParse(t *testing.T) {
	r := new(struct {
		Time ObsLocalTime `json:"time"`
	})
	if err := json.Unmarshal(mockObsTime, r); err != nil {
		t.Fatal(err)
	}
	if _, offset := time.Time(r.Time).Zone(); offset != -6*3600 {
		t.Fatalf("Offset is incorrect, got %d.", offset)
	}
}

func TestObsTimeParseString(t *testing.T) {
	r := new(ObsLocalTime)
	if err := r.setByParseString("2016-12-10 19:00:00-06:00"); err != nil {
		t.Fatal(err)
	}
}

func TestObsTimeParseV(t *testing.T) {
	r := new(ObsLocalTime)
	if err := r.setByParseV(1481396400, "-06:00"); err != nil {
		t.Fatal(err)
	}
}
