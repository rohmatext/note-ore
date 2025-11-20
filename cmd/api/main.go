package main

import (
	"errors"
	"net/http"
	"rohmatext/ore-note/internal/boostrap"
	"rohmatext/ore-note/internal/config"
	"rohmatext/ore-note/internal/server"
	"time"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db, err := config.NewDatabase(viperConfig, log)

	app := config.NewEcho(viperConfig)
	app.Validator = config.NewValidator()

	if err != nil {
		log.Fatalf("database init failed: %v", err)
	}

	srv := server.NewServer(&boostrap.BootstrapConfig{
		App:    app,
		DB:     db,
		Log:    log,
		Config: viperConfig,
	})

	go func() {
		if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("http server failed: %v", err)
		}
	}()

	srv.GracefulShutdown(5 * time.Second)
}
