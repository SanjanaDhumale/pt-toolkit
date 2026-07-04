package config

import (
	"log"
	"os"
	"path/filepath"
)

var AppConfig Config

func Init() {

	file := filepath.Join("config", "config.yaml")

	// First run: create default configuration
	if _, err := os.Stat(file); os.IsNotExist(err) {

		cfg := Default()

		if err := Save(cfg); err != nil {
			log.Fatal(err)
		}

		AppConfig = cfg
		return
	}

	// Existing installation: load configuration
	cfg, err := Load()
	if err != nil {
		log.Fatal(err)
	}

	AppConfig = cfg
}