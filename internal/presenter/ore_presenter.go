package presenter

import (
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
)

func OreSuccessResponse(data *entity.Ore) *ApiResponse[model.OreResponse] {
	user := model.OreResponse{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.OreResponse]{
		Message: "Ore retrived successfully",
		Data:    user,
	}
}

func OresSuccessResponse(data []*entity.Ore) *ApiResponse[[]model.OreResponse] {
	ores := make([]model.OreResponse, len(data))

	for i, ore := range data {
		ores[i] = model.OreResponse{
			ID:        ore.ID,
			Name:      ore.Name,
			CreatedAt: ore.CreatedAt,
			UpdatedAt: ore.UpdatedAt,
		}
	}

	return &ApiResponse[[]model.OreResponse]{
		Message: "Ores retrived successfully",
		Data:    ores,
	}
}

func CreateOreSuccessResponse(data *entity.Ore) *ApiResponse[model.OreResponse] {
	user := model.OreResponse{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.OreResponse]{
		Message: "Ore created successfully",
		Data:    user,
	}
}

func UpdateOreSuccessResponse(data *entity.Ore) *ApiResponse[model.OreResponse] {
	user := model.OreResponse{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.OreResponse]{
		Message: "Ore updated successfully",
		Data:    user,
	}
}
