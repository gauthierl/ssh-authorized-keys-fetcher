package main

import "testing"
import "time"

func TestDurationUnmarshalJSON(t *testing.T) {
	var d Duration
	var err error

	err = d.UnmarshalJSON([]byte(`"24h"`))
	if err != nil {
		t.Error(`"24h" should not produce error.`)
	}
	if d != Duration(24*time.Hour) {
		t.Error(`Cannot convert "24h" to Duration...`)
	}

	err = d.UnmarshalJSON([]byte(`"1"`))
	if err == nil {
		t.Error("Should not be able to convert 1 to Duration...")
	}
}
