package main

import (
	"testing"
	"time"
)

func TestConfigLoad(t *testing.T) {
	var err error

	config := NewConfig()

	err = config.Load("fixtures/config_not_found.json")
	if err == nil {
		t.Error("Load should return an error.")
	}

	err = config.Load("fixtures/config_invalid.json")
	if err == nil {
		t.Error("Load should return an error.")
	}

	err = config.Load("fixtures/config.json")
	if err != nil {
		t.Error("fixtures/config.json should be valid.")
	}

	if config.CachePath != "/var/lib/db/ssh_authorized_keys.json" {
		t.Errorf("CachePath should be equal to /var/lib/db/ssh_authorized_keys.json (%v).\n", config.CachePath)
	}

	if config.CacheTTL != Duration(24*time.Hour) {
		t.Errorf("CacheTTL should be equal to 24 hours (%v).\n", config.CacheTTL)
	}
}
