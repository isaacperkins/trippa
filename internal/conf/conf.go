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
	DB struct {
		Flavor   string `yaml:"flavor"`
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
		User     string `yaml:"user""`
		Pwd      string `yaml:"pwd"`
	} `yaml:"db"`
}

func New() *Config {
	var conf *Config

	yamlFile, err := os.ReadFile("conf/config.yaml")
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
