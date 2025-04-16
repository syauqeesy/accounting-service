package configuration

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	HTTP struct {
		Port string `json:"port"`
	} `json:"http"`
	GRPC struct {
		Port    string `json:"port"`
		Service struct {
			User string `json:"user"`
		} `json:"service"`
	} `json:"grpc"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
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
