package handler

import (
	"net/http"
	"rohmatext/ore-note/internal/delivery/http/validator"
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
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
	cursor := ctx.QueryParam("cursor")
	users, nextCursor, err := h.UserUseCase.GetUsersPaginated(ctx.Request().Context(), 20, cursor)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, presenter.UsersSuccessResponse(users, nextCursor))
}

func (h *UserHandler) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	user, err := h.UserUseCase.GetUserById(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, presenter.UserSuccessResponse(user))
}

func (h *UserHandler) Me(ctx echo.Context) (err error) {
	user := ctx.Get("auth").(*entity.User)
	return ctx.JSON(http.StatusOK, presenter.UserSuccessResponse(user))
}

func (h *UserHandler) Store(ctx echo.Context) (err error) {
	request := new(model.CreateUserRequest)
	if err = ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctx.Validate(request); err != nil {
		h.Log.Warnf("invalid request body: %+v", err)
		if ve, ok := err.(*validator.ValidationError); ok {
			return ctx.JSON(http.StatusUnprocessableEntity, ve)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user, err := h.UserUseCase.CreateOperator(ctx.Request().Context(), request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, presenter.CreateUserSuccessResponse(user))
}

func (h *UserHandler) Update(ctx echo.Context) (err error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	data, err := h.UserUseCase.GetUserById(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	request := new(model.UpdateUserRequest)
	request.ID = data.ID
	if err = ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctx.Validate(request); err != nil {
		h.Log.Warnf("invalid request body: %+v", err)
		if ve, ok := err.(*validator.ValidationError); ok {
			return ctx.JSON(http.StatusUnprocessableEntity, ve)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user, err := h.UserUseCase.UpdateUser(ctx.Request().Context(), request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, presenter.UpdateUserSuccessResponse(user))
}

func (h *UserHandler) Delete(ctx echo.Context) (err error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	user, err := h.UserUseCase.GetUserById(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	auth := ctx.Get("auth").(*entity.User)
	if user.ID == auth.ID {
		return echo.NewHTTPError(http.StatusForbidden, "Self-deletion is not allowed.")
	}

	if err := h.UserUseCase.DeleteUser(ctx.Request().Context(), uint(id)); err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusNoContent, "")
}
