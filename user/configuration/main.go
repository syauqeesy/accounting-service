package configuration

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	HTTP struct {
		Port string `json:"port"`
	}
}

func Load(path string) (*Configuration, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configuration Configuration

	err = json.Unmarshal(bytes, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
