package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const gatorConfig = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	bytes, err := os.ReadFile(path)

	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(bytes, &config); err != nil {
		return config, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return home, err
	}

	return filepath.Join(home, gatorConfig), nil
}
