package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CookieService struct{}

func (cs *CookieService) SetRefreshToken(ctx echo.Context, token string, expires time.Time) {
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Expires:  expires,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	ctx.SetCookie(cookie)
}

func (cs *CookieService) GetRefreshToken(ctx echo.Context) (*string, error) {
	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		return nil, err
	}
	return &cookie.Value, nil
}
