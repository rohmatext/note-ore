package boostrap

import (
	dHttp "rohmatext/ore-note/internal/delivery/http"
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
	cookieService := &dHttp.CookieService{}

	roleRepository := repository.NewRoleRepository(cfg.Log)
	userRepository := repository.NewUserRepository(cfg.Log)
	refreshTokenRepository := repository.NewRefreshTokenRepository(cfg.Log)

	roleUseCase := usecase.NewRoleUseCase(cfg.DB, cfg.Log, roleRepository)
	userUseCase := usecase.NewUserUseCase(cfg.DB, cfg.Log, refreshTokenRepository, userRepository, roleRepository, tokenService)
	refreshTokenUseCase := usecase.NewRefreshTokenUseCase(cfg.DB, cfg.Log, refreshTokenRepository)

	authHandler := handler.NewAuthHandler(cfg.Log, cookieService, userUseCase, refreshTokenUseCase)
	userHandler := handler.NewUserHandler(cfg.Log, userUseCase)
	roleHandler := handler.NewRoleHandler(cfg.Log, roleUseCase)

	routeConfig := route.RouteConfig{
		App:            cfg.App,
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		RoleHandler:    roleHandler,
		AuthMiddleware: middleware.AuthMiddleware(userUseCase, cfg.Config),
		RoleMiddleware: middleware.RoleMiddleware,
	}
	routeConfig.SetupRoutes()

	return cfg.App
}
