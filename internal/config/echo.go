package config

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func NewEcho(config *viper.Viper) *echo.Echo {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	allowOrigins := strings.Split(config.GetString("FRONTEND_URLS"), ",")
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowCredentials: true,
	}))

	return app
}
