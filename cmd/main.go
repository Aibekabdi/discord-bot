package main

import (
	"log"
	"mybot/internal/app"
	"mybot/internal/models"
	"mybot/pkg"
)

func main() {
	conf, err := pkg.NewConfig(models.ConfigPath)
	if err != nil {
		log.Fatalf("error occured while getting configs from file: %v", err)
	}
	app.Run(conf)

}
