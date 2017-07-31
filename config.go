package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config structure
type Config struct {
	CachePath string   `json:"cache_path"`
	CacheTTL  Duration `json:"cache_ttl"`
	FetchURL  string   `json:"fetch_url"`
}

// NewConfig constructor
func NewConfig() *Config {
	return &Config{}
}

// Load the config from a JSON file
func (config *Config) Load(filePath string) (err error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return
	}

	return
}
