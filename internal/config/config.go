package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	strHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	strConfigFile := filepath.Join(strHomeDir, ".gatorconfig.json")

	return strConfigFile, nil
}

func Read() (Config, error) {
	var configResp Config

	strConfigFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(strConfigFile)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, &configResp)
	if err != nil {
		return Config{}, err
	}

	return configResp, nil
}

func (c *Config) SetUser(name string) error {
	// Put the provided parameter into the Config struct
	c.CurrentUserName = name

	strConfigFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Marshal the struct into a []byte slice
	jsonConfig, err := json.Marshal(*c)
	if err != nil {
		return err
	}

	err = os.WriteFile(strConfigFile, jsonConfig, 0644)
	if err != nil {
		return err
	}

	return nil
}
