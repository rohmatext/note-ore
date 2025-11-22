package handler

import (
	"net/http"
	"rohmatext/ore-note/internal/presenter"
	"rohmatext/ore-note/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type RoleHandler struct {
	RoleUseCase usecase.RoleUseCase
	Log         *logrus.Logger
}

func NewRoleHandler(log *logrus.Logger, roleUC usecase.RoleUseCase) *RoleHandler {
	return &RoleHandler{
		RoleUseCase: roleUC,
		Log:         log,
	}
}

func (h *RoleHandler) List(ctx echo.Context) error {
	roles, err := h.RoleUseCase.GetAllRoles(ctx.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, presenter.RolesSuccessResponse(roles))
}
