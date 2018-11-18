package aqicn

import (
	"encoding/json"
	"testing"
)

func TestNumberInt(t *testing.T) {
	d := new(struct {
		N Number `json:"n"`
	})
	if err := json.Unmarshal([]byte(`{"n":1}`), d); err != nil {
		t.Fatal(err)
	} else if d.N.Int() != 1 {
		t.Fatalf("Number is incorrect: got %d.", d.N.Int())
	} else if !d.N.IsInt() {
		t.Fatal("Number should be int")
	}
}

func TestNumberFloat(t *testing.T) {
	d := new(struct {
		N Number `json:"n"`
	})
	if err := json.Unmarshal([]byte(`{"n":1.7}`), d); err != nil {
		t.Fatal(err)
	} else if d.N.Int() != 1 {
		t.Fatalf("Number is incorrect: got %d.", d.N.Int())
	} else if !d.N.IsFloat() {
		t.Fatal("Number should be float")
	}
}
