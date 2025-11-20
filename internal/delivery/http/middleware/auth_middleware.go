package middleware

import (
	"net/http"
	"rohmatext/ore-note/internal/usecase"

	jwtService "rohmatext/ore-note/internal/infrastructure/jwt"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func AuthMiddleware(userUC usecase.UserUseCase, config *viper.Viper) echo.MiddlewareFunc {
	jwtMw := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.GetString("JWT_SECRET")),
		ErrorHandler: func(ctx echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token.").SetInternal(err)
		},
		SigningMethod: jwt.SigningMethodHS256.Alg(),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtService.MapClaims)
		},
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return jwtMw(func(ctx echo.Context) (err error) {
			userToken, ok := ctx.Get("user").(*jwt.Token)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invlid or expired token").SetInternal(err)
			}

			claims, ok := userToken.Claims.(*jwtService.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invlid or expired token").SetInternal(err)
			}

			user, err := userUC.GetUser(ctx.Request().Context(), claims.UserId)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invlid or expired token").SetInternal(err)
			}

			ctx.Set("auth", user)
			return next(ctx)
		})
	}

}
