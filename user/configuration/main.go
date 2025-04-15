package configuration

import (
	"encoding/json"
	"os"
)

type Config struct {
	HTTP struct {
		Port string `json:"port"`
	}
}

func Load(path string) (*Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
