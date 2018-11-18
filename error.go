package aqicn

import (
	"errors"
)

var (
	// ErrOverQuota means that the request is over quota limits
	ErrOverQuota = errors.New("the request is over quota limits")
	// ErrInvalidKey means that the key is not valid
	ErrInvalidKey = errors.New("the key is not valid")
	// ErrUnknownCity means that the city is unknown
	ErrUnknownCity = errors.New("the city is unknown")
)
