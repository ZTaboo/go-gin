package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var (
	Port string
)

func init() {
	var config Config
	file, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Print(err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Print(err)
	}
	Port = ":" + config.Host.Port
}
