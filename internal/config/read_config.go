package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() Config {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("error getting config file path: %v\n", err)
		return Config{}
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return Config{}
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Printf("Error unmarshaling config file: %v\n", err)
		return Config{}
	}

	return cfg
}
