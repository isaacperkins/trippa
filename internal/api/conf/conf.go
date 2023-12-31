package conf

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var Conf *Config

type Config struct {
	Http struct {
		Schema string `yaml:"scheme"`
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
	} `yaml:"http"`
}

func New(fp string) *Config {
	var conf *Config

	yamlFile, err := os.ReadFile(fp)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	Conf = conf
	return conf
}

func GetConfig() *Config {
	return Conf
}
