package handler

import (
	"errors"
	"net/http"
	cookie "rohmatext/ore-note/internal/delivery/http"
	"rohmatext/ore-note/internal/delivery/http/validator"
	"rohmatext/ore-note/internal/model"
	"rohmatext/ore-note/internal/presenter"
	"rohmatext/ore-note/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	Log                 *logrus.Logger
	Cookie              *cookie.CookieService
	UserUseCase         usecase.UserUseCase
	RefreshTokenUseCase usecase.RefreshTokenUseCase
}

func NewAuthHandler(log *logrus.Logger, cookie *cookie.CookieService, user usecase.UserUseCase, refreshToken usecase.RefreshTokenUseCase) *AuthHandler {
	return &AuthHandler{
		Log:                 log,
		Cookie:              cookie,
		UserUseCase:         user,
		RefreshTokenUseCase: refreshToken,
	}
}

func (h *AuthHandler) Login(ctx echo.Context) (err error) {
	req := new(model.LoginUserRequest)
	if err = ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctx.Validate(req); err != nil {
		h.Log.Warnf("invalid request body: %+v", err)
		if ve, ok := err.(*validator.ValidationError); ok {
			return ctx.JSON(http.StatusUnprocessableEntity, ve)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := h.UserUseCase.Login(ctx.Request().Context(), req)
	if err != nil {
		h.Log.Warnf("failed to login user: %+v", err)
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			return ctx.JSON(http.StatusUnprocessableEntity, validator.ValidationError{
				Message: "Validation failed",
				Errors: map[string]string{
					"username": "These credentials do not match our records.",
				},
			})
		}

		return err
	}

	h.Cookie.SetRefreshToken(ctx, token.RefreshToken.Token, token.RefreshToken.ExpiresAt)

	return ctx.JSON(http.StatusOK, presenter.UserLoginSuccessResponse(token.AccessToken))
}

func (h *AuthHandler) RefreshToken(ctx echo.Context) error {
	tokenStr, err := h.Cookie.GetRefreshToken(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := h.UserUseCase.RefreshAccessToken(ctx.Request().Context(), *tokenStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	h.Cookie.SetRefreshToken(ctx, token.RefreshToken.Token, token.RefreshToken.ExpiresAt)

	return ctx.JSON(http.StatusOK, presenter.UserLoginSuccessResponse(token.AccessToken))
}
