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
	RoleHandler    *handler.RoleHandler
	OreHandler     *handler.OreHandler
	SourceHandler  *handler.SourceHandler
	AuthMiddleware echo.MiddlewareFunc
	RoleMiddleware func(...string) echo.MiddlewareFunc
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

	auth.GET("/api/me", r.UserHandler.Me)

	auth.GET("/api/roles", r.RoleHandler.List)

	auth.GET("/api/users", r.UserHandler.List, r.RoleMiddleware("admin"))
	auth.GET("/api/users/:id", r.UserHandler.Get, r.RoleMiddleware("admin"))
	auth.POST("/api/users", r.UserHandler.Store, r.RoleMiddleware("admin"))
	auth.PATCH("/api/users/:id", r.UserHandler.Update, r.RoleMiddleware("admin"))
	auth.DELETE("/api/users/:id", r.UserHandler.Delete, r.RoleMiddleware("admin"))

	auth.GET("/api/ores", r.OreHandler.List, r.RoleMiddleware("admin"))
	auth.GET("/api/ores/:id", r.OreHandler.Show, r.RoleMiddleware("admin"))
	auth.POST("/api/ores", r.OreHandler.Store, r.RoleMiddleware("admin"))
	auth.PATCH("/api/ores/:id", r.OreHandler.Update, r.RoleMiddleware("admin"))
	auth.DELETE("/api/ores/:id", r.OreHandler.Delete, r.RoleMiddleware("admin"))

	auth.GET("/api/sources", r.SourceHandler.List)
	auth.GET("/api/sources/:id", r.SourceHandler.Show)
	auth.POST("/api/sources", r.SourceHandler.Store)
	auth.PATCH("/api/sources/:id", r.SourceHandler.Update)
	auth.DELETE("/api/sources/:id", r.SourceHandler.Delete)
}

func upHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "HTTP request received.",
	})
}
