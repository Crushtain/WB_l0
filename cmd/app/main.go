package main

import (
	"WB_L0/config"
	"WB_L0/internal/app"
	"log"
)

func main() {
	configPath := "config/config.toml"
	cfg, err := config.ParseTOML(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	app.Run(cfg)
}

//func findConfigPath()
