package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	localConfigPath = "config/config.yaml"
	envConfigPath   = "config/config-docker.yaml"
)

type Conf struct {
	Http     map[string]string `yaml:"http"`
	Database map[string]string `yaml:"database"`
}

func (c *Conf) GetConf() *Conf {
	config := localConfigPath
	if isRunningInDockerContainer() {
		config = envConfigPath
	}

	yamlFile, err := os.ReadFile(config)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func isRunningInDockerContainer() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
