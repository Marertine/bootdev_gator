package config

import (
	"encoding/json"
	"errors"
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
	if errors.Is(err, os.ErrNotExist) {
		// Config doesn't yet exist, return default Config and no error
		return configResp, nil
	}
	if err != nil {
		// All other errors
		return Config{}, err
	}

	err = json.Unmarshal(data, &configResp)
	if err != nil {
		return Config{}, err
	}

	return configResp, nil
}

func (c *Config) SetUser(name string) error {
	// If this were func(c Config) it would be working with a COPY
	// of the config, but we want to actually amend the original so
	// we must work with a pointer to it!!

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
