package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"log"
)

type Config struct {
	App    App        `toml:"App"`
	DB     DB         `toml:"Database"`
	Server HttpServer `toml:"HttpServer"`
	logger *logrus.Logger
}

type App struct {
	Name    string `toml:"Name"`
	Version string `toml:"Version"`
}

type DB struct {
	Name        string `toml:"Name"`
	Host        string `toml:"Host"`
	Port        string `toml:"Port"`
	Schema      string `toml:"Schema"`
	User        string `env:"User"`
	Password    string `env:"Password"`
	MaxPoolSize int    `toml:"MaxPoolSize"`
}

type HttpServer struct {
	bindAddr string     `toml:"bindAddr"`
	logLevel log.Logger `toml:"logLevel"`
}

func ParseTOML(path string) (*Config, error) {
	var configData Config
	_, err := toml.Decode(path, &configData)
	if err != nil {
		return nil, err
	}
	return &configData, nil
}

// func ParseConfig()
