package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func (c *Config) GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("servers.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)

	}
	err = yaml.Unmarshal([]byte(yamlFile), &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)

	}
	if err != nil {
		log.Fatalf("error: %v", err)

	}
	return c
}
