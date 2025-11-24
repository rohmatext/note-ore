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

type SourceHandler struct {
	SourceUseCase usecase.SourceUseCase
	Log           *logrus.Logger
}

func NewSourceHandler(log *logrus.Logger, sourceUC usecase.SourceUseCase) *SourceHandler {
	return &SourceHandler{
		SourceUseCase: sourceUC,
		Log:           log,
	}
}

func (h *SourceHandler) List(ctx echo.Context) error {
	sources, err := h.SourceUseCase.GetAllSources(ctx.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, presenter.SourcesSuccessResponse(sources))
}

func (h *SourceHandler) Show(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	source, err := h.SourceUseCase.GetSourceById(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, presenter.SourceSuccessResponse(source))
}

func (h *SourceHandler) Store(ctx echo.Context) error {
	request := new(model.CreateSourceRequest)
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

	source, err := h.SourceUseCase.CreateSource(ctx.Request().Context(), request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, presenter.CreateSourceSuccessResponse(source))
}

func (h *SourceHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	data, err := h.SourceUseCase.GetSourceById(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	request := new(model.UpdateSourceRequest)
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

	source, err := h.SourceUseCase.UpdateSource(ctx.Request().Context(), request, data.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, presenter.UpdateSourceSuccessResponse(source))
}

func (h *SourceHandler) Delete(ctx echo.Context) (err error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	source, err := h.SourceUseCase.GetSourceById(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := h.SourceUseCase.DeleteSource(ctx.Request().Context(), source.ID); err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusNoContent, "")
}
