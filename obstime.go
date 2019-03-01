package aqicn

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// ObsLocalTime represents the observed time of the data
type ObsLocalTime time.Time

func (c *ObsLocalTime) setByParseString(raw string) error {
	t, err := time.Parse("2006-01-02 15:04:05-07:00", raw)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05-0700", raw)
		if err != nil {
			return err
		}
	}
	*c = ObsLocalTime(t)
	return nil
}

func (c *ObsLocalTime) setByParseV(v int64, tz string) error {
	if len(tz) != 6 {
		return errors.New("tz offset length is incorrect")
	}
	offseth, err := strconv.Atoi(tz[1:3])
	if err != nil {
		return err
	}
	offsetm, err := strconv.Atoi(tz[4:6])
	if err != nil {
		return err
	}
	var offsetmult int
	switch tz[0] {
	case '+':
		offsetmult = 1
	case '-':
		offsetmult = -1
	default:
		return errors.New("tz offset sign incorrect")
	}
	*c = ObsLocalTime(time.Unix(v+int64(offsetmult*(offseth*3600+offsetm*60)), 0))
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (c *ObsLocalTime) UnmarshalJSON(b []byte) error {
	raw := new(struct {
		S          string `json:"s"`
		TZ         string `json:"tz"`
		V          int64  `json:"v"`
		SAlternate string `json:"stime"`
		VAlternate int64  `json:"vtime"`
	})
	if err := json.Unmarshal(b, raw); err != nil {
		return err
	}
	if raw.SAlternate != "" && raw.S == "" {
		raw.S = raw.SAlternate
	}
	if raw.VAlternate != 0 && raw.V == 0 {
		raw.V = raw.VAlternate
	}
	switch {
	case c.setByParseString(raw.S+raw.TZ) == nil:
		return nil
	case c.setByParseV(raw.V, raw.TZ) == nil:
		return nil
	default:
		return errors.New("Failed to parse obs time")
	}
}

func (c ObsLocalTime) String() string {
	return time.Time(c).String()
}
