package config

import "log"

var AppConfig *Config

func Init() {
	var err error

	AppConfig, err = LoadConfig("configs/toolkit.yaml")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
}