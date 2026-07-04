package config

import "log"

var AppConfig Config

func Init() {

	cfg, err := Load()

	if err != nil {

		log.Fatal(err)
	}

	AppConfig = cfg
}