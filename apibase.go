package aqicn

import (
	"encoding/json"
	"errors"
)

type apiBaseResponse []byte

func getData(d []byte) (apiBaseResponse, error) {
	r := new(struct {
		Status string          `json:"status"`
		Msg    string          `json:"message"`
		Data   json.RawMessage `json:"data"`
	})
	if err := json.Unmarshal(d, r); err != nil {
		return nil, err
	}
	if r.Status != "ok" {
		var err error
		switch r.Msg {
		case "Over quota":
			err = ErrOverQuota
		case "Invalid key":
			err = ErrInvalidKey
		case "Unknown City":
			err = ErrUnknownCity
		default:
			err = errors.New(r.Msg)
		}
		return nil, err
	}
	return []byte(r.Data), nil
}
