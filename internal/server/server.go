package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"rohmatext/ore-note/internal/boostrap"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Server *http.Server
	App    *echo.Echo
	Log    *logrus.Logger
}

func NewServer(config *boostrap.BootstrapConfig) *ServerConfig {
	e := config.Bootstrap()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Config.GetInt("PORT")),
		Handler:      e,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &ServerConfig{
		Server: srv,
		App:    e,
		Log:    config.Log,
	}
}

func (s *ServerConfig) Run() error {
	return s.App.StartServer(s.Server)
}

func (s *ServerConfig) GracefulShutdown(timeout time.Duration) {
	// listen for OS signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	s.Log.Info("Graceful shutdown triggered")

	shudownContext, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// stop accepting new request & close idle connections
	if err := s.App.Shutdown(shudownContext); err != nil {
		s.Log.Errorf("Echo shutdown error: %+v", err)
	}

	if err := s.Server.Shutdown(shudownContext); err != nil {
		s.Log.Errorf("Server shutdown error: %+v", err)
		_ = s.Server.Close()
	}

	s.Log.Info("Server exited gracefully")

}
