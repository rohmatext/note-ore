package presenter

import (
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
)

func RolesSuccessResponse(data []*entity.Role) *ApiResponse[[]model.RoleResponse] {
	roles := make([]model.RoleResponse, len(data))

	for i, role := range data {
		roles[i] = model.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
		}
	}

	return &ApiResponse[[]model.RoleResponse]{
		Message: "Roles retrived successfully",
		Data:    roles,
	}
}
