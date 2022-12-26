package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Env      map[string]string `yaml:"env"`
	Http     map[string]string `yaml:"http"`
	Database map[string]string `yaml:"database"`
}

func (c *Conf) GetConfig() *Conf {
	if isRunningInDockerContainer() {
		dockerConfig(c)
	} else {
		localConfig(c)
	}

	return c
}

func dockerConfig(c *Conf) {
	c.Env = make(map[string]string)
	c.Env["name"] = "docker"

	c.Http = make(map[string]string)
	c.Http["port"] = os.Getenv("HTTP_PORT")

	c.Database = make(map[string]string)
	c.Database["user"] = os.Getenv("DB_USER")
	c.Database["pass"] = os.Getenv("DB_PASS")
	c.Database["name"] = os.Getenv("DB_NAME")
	c.Database["port"] = os.Getenv("DB_PORT")
}

func localConfig(c *Conf) {
	localConfigPath := "./config-local.yaml"

	yamlFile, err := os.ReadFile(localConfigPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

}

func isRunningInDockerContainer() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
