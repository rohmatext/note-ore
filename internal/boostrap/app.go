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
	oreRepository := repository.NewOreRepository(cfg.Log)
	sourceRepository := repository.NewSourceRepository(cfg.Log)
	productionRepository := repository.NewProductionRepository(cfg.Log)

	roleUseCase := usecase.NewRoleUseCase(cfg.DB, cfg.Log, roleRepository)
	userUseCase := usecase.NewUserUseCase(cfg.DB, cfg.Log, refreshTokenRepository, userRepository, roleRepository, tokenService)
	refreshTokenUseCase := usecase.NewRefreshTokenUseCase(cfg.DB, cfg.Log, refreshTokenRepository)
	oreUseCase := usecase.NewOreUseCase(cfg.DB, cfg.Log, oreRepository)
	sourceUseCase := usecase.NewSourceUseCase(cfg.DB, cfg.Log, sourceRepository)
	productionUseCase := usecase.NewProductionUseCase(cfg.DB, cfg.Log, productionRepository, userRepository, oreRepository, sourceRepository)

	authHandler := handler.NewAuthHandler(cfg.Log, cookieService, userUseCase, refreshTokenUseCase)
	userHandler := handler.NewUserHandler(cfg.Log, userUseCase)
	roleHandler := handler.NewRoleHandler(cfg.Log, roleUseCase)
	oreHandler := handler.NewOreHandler(cfg.Log, oreUseCase)
	sourceHandler := handler.NewSourceHandler(cfg.Log, sourceUseCase)
	productionHandler := handler.NewProductionHandler(cfg.Log, productionUseCase)

	routeConfig := route.RouteConfig{
		App:               cfg.App,
		AuthHandler:       authHandler,
		UserHandler:       userHandler,
		RoleHandler:       roleHandler,
		OreHandler:        oreHandler,
		SourceHandler:     sourceHandler,
		ProductionHandler: productionHandler,
		AuthMiddleware:    middleware.AuthMiddleware(userUseCase, cfg.Config),
		RoleMiddleware:    middleware.RoleMiddleware,
	}
	routeConfig.SetupRoutes()

	return cfg.App
}
