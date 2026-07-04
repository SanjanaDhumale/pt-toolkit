package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Load() (Config, error) {

	var cfg Config

	file := filepath.Join("config", "config.yaml")

	data, err := os.ReadFile(file)

	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)

	return cfg, err
}