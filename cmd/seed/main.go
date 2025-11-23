package main

import (
	"rohmatext/ore-note/internal/config"
	"rohmatext/ore-note/internal/seeder"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db, err := config.NewDatabase(viperConfig, log)

	if err != nil {
		panic("failed to connect database")
	}

	seeder.Run(db)
}
