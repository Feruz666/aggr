package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Rds      RedisConfig    `yaml:"redis"`
	LogLevel string         `yaml:"logLevel"`
}

type ServerConfig struct {
	ServerAddr string `yaml:"serverAddr"`
}

type DatabaseConfig struct {
	Url string `yaml:"url"`
}

type RedisConfig struct {
	RedisAddr string `yaml:"redisAddr"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	f, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg, nil
}
