package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg Config) SetUser(userName string) {
	cfg.CurrentUserName = userName
	write(cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting user home directory: %w", err)
	}
	configFilePath := homeDir + "/" + configFileName
	return configFilePath, nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %w", err)
	}

	data, err := json.Marshal(&cfg)
	if err != nil {
		return fmt.Errorf("error marshaling config file: %w", err)
	}

	if err := os.WriteFile(configFilePath, data, 0666); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}
