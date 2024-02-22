package main

import (
	"WB_l0/config"
	"WB_l0/internal/app"
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
