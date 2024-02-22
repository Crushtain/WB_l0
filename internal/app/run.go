package app

import (
	"WB_l0/config"
	"WB_l0/pkg/postgres"
	"context"
	"log"
)

func Run(config *config.Config) {
	connString := postgres.GetConnString(&config.DB)
	pg, err := postgres.New(connString)
	if err != nil {
		log.Fatalln("Error connection to database")
	}
	// defer pg.Pool.Close() // better to make a check if pool != nil
	defer pg.Close()

	err = pg.Pool.Ping(context.Background()) //context if it would be in need to make ping-time-check
	if err != nil {
		log.Fatalln("Error ping database")
	}

}
