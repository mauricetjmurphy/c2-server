package config

import (
	"encoding/json"
	"os"
	"time"
)

// Config holds the configuration settings for the implant
type Config struct {
	Host     string        `json:"host"`
	Port     string        `json:"port"`
	URI      string        `json:"uri"`
	Timeout  time.Duration `json:"timeout"`
	Interval time.Duration `json:"interval"`
}

// LoadConfig loads configuration from a file
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := &Config{}
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}

	// Convert timeout and interval from seconds to duration
	cfg.Timeout *= time.Second
	cfg.Interval *= time.Second

	return cfg, nil
}
