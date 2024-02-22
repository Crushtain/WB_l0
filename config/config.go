package config

import "github.com/BurntSushi/toml"

type Config struct {
}

func GetCongigFromToml(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
