package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Session string `yaml:"session"`
	} `yaml:"app"`
}

func New(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
