package boostrap

import (
	"rohmatext/ore-note/internal/delivery/http/handler"
	"rohmatext/ore-note/internal/delivery/http/middleware"
	"rohmatext/ore-note/internal/delivery/http/route"
	"rohmatext/ore-note/internal/infrastructure/jwt"
	"rohmatext/ore-note/internal/repository"
	"rohmatext/ore-note/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App    *echo.Echo
	DB     *gorm.DB
	Log    *logrus.Logger
	Config *viper.Viper
}

func (cfg *BootstrapConfig) Bootstrap() *echo.Echo {
	tokenService := jwt.NewJWTService(cfg.Config.GetString("JWT_SECRET"))

	authRepository := repository.NewUserRepository(cfg.Log)

	userUseCase := usecase.NewUserUseCase(cfg.DB, cfg.Log, authRepository)

	authHandler := handler.NewAuthHandler(cfg.Log, userUseCase, tokenService)
	userHandler := handler.NewUserHandler(cfg.Log, userUseCase)

	routeConfig := route.RouteConfig{
		App:            cfg.App,
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		AuthMiddleware: middleware.AuthMiddleware(userUseCase, cfg.Config),
	}
	routeConfig.SetupRoutes()

	return cfg.App
}
