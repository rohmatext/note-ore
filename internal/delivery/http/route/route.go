package route

import (
	"net/http"
	"rohmatext/ore-note/internal/delivery/http/handler"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	App            *echo.Echo
	AuthHandler    *handler.AuthHandler
	UserHandler    *handler.UserHandler
	AuthMiddleware echo.MiddlewareFunc
}

func (r *RouteConfig) SetupRoutes() {
	r.SetupGuestRoutes()
	r.SetupAuthRoutes()
}

func (r *RouteConfig) SetupGuestRoutes() {
	r.App.GET("/api/up", upHandler)
	r.App.POST("/api/login", r.AuthHandler.Login)
	r.App.POST("/api/refresh", r.AuthHandler.RefreshToken)
}

func (r *RouteConfig) SetupAuthRoutes() {
	auth := r.App.Group("", r.AuthMiddleware)

	auth.GET("/api/me", r.AuthHandler.Me)

	auth.GET("/api/users", r.UserHandler.List)
	auth.GET("/api/users/:id", r.UserHandler.Get)
}

func upHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "HTTP request received.",
	})
}
