package aqicn

import (
	"encoding/json"
	"strconv"
)

// Number represents a JS number which could either be an int or float32
type Number struct {
	nint   *int
	nfloat *float32
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (c *Number) UnmarshalJSON(d []byte) error {
	i := new(int)
	if err := json.Unmarshal(d, i); err == nil {
		c.nint = i
		c.nfloat = nil
		return nil
	} else if _, ok := err.(*json.UnmarshalTypeError); ok {
		f := new(float32)
		if err := json.Unmarshal(d, f); err != nil {
			return err
		}
		c.nint = nil
		c.nfloat = f
		return nil
	} else {
		return err
	}
}

// Int converts the number to int. Panics if the number is not properly initialized.
func (c Number) Int() int {
	switch {
	case c.nint != nil:
		return *c.nint
	case c.nfloat != nil:
		return int(*c.nfloat)
	default:
		panic("Not initialized")
	}
}

// Float converts the number to float32. Panics if the number is not properly initialized.
func (c Number) Float() float32 {
	switch {
	case c.nint != nil:
		return float32(*c.nint)
	case c.nfloat != nil:
		return *c.nfloat
	default:
		panic("Not initialized")
	}
}

// IsInt returns if the underlying number is an int
func (c Number) IsInt() bool {
	return c.nint != nil
}

// IsFloat returns if the underlying number is a float
func (c Number) IsFloat() bool {
	return c.nfloat != nil
}

func (c Number) String() string {
	switch {
	case c.nint != nil:
		return strconv.Itoa(*c.nint)
	case c.nfloat != nil:
		return strconv.FormatFloat(float64(*c.nfloat), 'f', -1, 32)
	default:
		return "<Number> not initialized"
	}
}
