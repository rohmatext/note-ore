package handler

import (
	"net/http"
	"rohmatext/ore-note/internal/delivery/http/validator"
	"rohmatext/ore-note/internal/model"
	"rohmatext/ore-note/internal/presenter"
	"rohmatext/ore-note/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type OreHandler struct {
	OreUseCase usecase.OreUseCase
	Log        *logrus.Logger
}

func NewOreHandler(log *logrus.Logger, oreUC usecase.OreUseCase) *OreHandler {
	return &OreHandler{
		OreUseCase: oreUC,
		Log:        log,
	}
}

func (h *OreHandler) List(ctx echo.Context) error {
	ores, err := h.OreUseCase.GetAllOres(ctx.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, presenter.OresSuccessResponse(ores))
}

func (h *OreHandler) Show(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	ore, err := h.OreUseCase.GetOreById(ctx.Request().Context(), uint16(id))
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, presenter.OreSuccessResponse(ore))
}

func (h *OreHandler) Store(ctx echo.Context) error {
	request := new(model.CreateOreRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(request); err != nil {
		h.Log.Warnf("invalid request body: %+v", err)
		if ve, ok := err.(*validator.ValidationError); ok {
			return ctx.JSON(http.StatusUnprocessableEntity, ve)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ore, err := h.OreUseCase.CreateOre(ctx.Request().Context(), request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, presenter.CreateOreSuccessResponse(ore))
}

func (h *OreHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	data, err := h.OreUseCase.GetOreById(ctx.Request().Context(), uint16(id))
	if err != nil {
		return echo.ErrNotFound
	}

	request := new(model.UpdateOreRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(request); err != nil {
		h.Log.Warnf("invalid request body: %+v", err)
		if ve, ok := err.(*validator.ValidationError); ok {
			return ctx.JSON(http.StatusUnprocessableEntity, ve)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ore, err := h.OreUseCase.UpdateOre(ctx.Request().Context(), request, data.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, presenter.UpdateOreSuccessResponse(ore))
}

func (h *OreHandler) Delete(ctx echo.Context) (err error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	ore, err := h.OreUseCase.GetOreById(ctx.Request().Context(), uint16(id))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := h.OreUseCase.DeleteOre(ctx.Request().Context(), ore.ID); err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusNoContent, "")
}
