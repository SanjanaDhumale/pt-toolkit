package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Save(cfg Config) error {

	if err := os.MkdirAll("config", 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)

	if err != nil {
		return err
	}

	file := filepath.Join("config", "config.yaml")

	return os.WriteFile(file, data, 0644)
}