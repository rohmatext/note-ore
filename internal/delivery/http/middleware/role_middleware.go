package middleware

import (
	"rohmatext/ore-note/internal/entity"
	"slices"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if !slices.Contains(roles, ctx.Get("auth").(*entity.User).Role.Name) {
				return echo.ErrUnauthorized
			}
			return next(ctx)
		}
	}
}
