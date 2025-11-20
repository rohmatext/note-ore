package handler

import (
	"net/http"
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/presenter"
	"rohmatext/ore-note/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
	Log         *logrus.Logger
}

func NewUserHandler(log *logrus.Logger, userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: userUseCase,
		Log:         log,
	}
}

func (h *UserHandler) List(ctx echo.Context) error {
	users, err := h.UserUseCase.GetUsers(ctx.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, presenter.UsersSuccessResponse(users))
}

func (h *UserHandler) Get(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrInternalServerError
	}

	user, err := h.UserUseCase.GetUser(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, presenter.UserSuccessResponse(user))
}

func (h *AuthHandler) Me(ctx echo.Context) (err error) {
	user := ctx.Get("auth").(*entity.User)
	return ctx.JSON(http.StatusOK, presenter.UserSuccessResponse(user))
}
