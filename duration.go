package main

import (
	"strconv"
	"time"
)

// Duration extend to implement UnmarshalJSON
type Duration time.Duration

// UnmarshalJSON implementation for Duration
func (d *Duration) UnmarshalJSON(data []byte) (err error) {
	s := string(data)

	s, err = strconv.Unquote(s)
	if err != nil {
		return
	}

	t, err := time.ParseDuration(s)
	if err != nil {
		return
	}

	*d = Duration(t)
	return
}
