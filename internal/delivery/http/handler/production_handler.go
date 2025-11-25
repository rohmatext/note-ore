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

type ProductionHandler struct {
	ProductionUseCase usecase.ProductionUseCase
	Log               *logrus.Logger
}

func NewProductionHandler(log *logrus.Logger, productionUC usecase.ProductionUseCase) *ProductionHandler {
	return &ProductionHandler{
		ProductionUseCase: productionUC,
		Log:               log,
	}
}

func (h *ProductionHandler) List(ctx echo.Context) error {
	productions, err := h.ProductionUseCase.GetAllProductions(ctx.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, presenter.ProductionsSuccessResponse(productions))
}

func (h *ProductionHandler) Show(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	production, err := h.ProductionUseCase.GetProductionById(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, presenter.ProductionSuccessResponse(production))
}

func (h *ProductionHandler) Store(ctx echo.Context) error {
	request := new(model.CreateProductionRequest)
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

	auth := ctx.Get("auth").(*entity.User)
	production, err := h.ProductionUseCase.CreateProduction(ctx.Request().Context(), auth.ID, request)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, presenter.CreateProductionSuccessResponse(production))
}

func (h *ProductionHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	data, err := h.ProductionUseCase.GetProductionById(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	request := new(model.UpdateProductionRequest)
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

	auth := ctx.Get("auth").(*entity.User)
	production, err := h.ProductionUseCase.UpdateProduction(ctx.Request().Context(), auth.ID, request, data.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, presenter.UpdateProductionSuccessResponse(production))
}

func (h *ProductionHandler) Delete(ctx echo.Context) (err error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	production, err := h.ProductionUseCase.GetProductionById(ctx.Request().Context(), uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := h.ProductionUseCase.DeleteProduction(ctx.Request().Context(), production.ID); err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusNoContent, "")
}
