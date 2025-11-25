package presenter

import (
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
)

func ProductionSuccessResponse(data *entity.Production) *ApiResponse[model.ProductionResponse] {
	production := model.ProductionResponse{
		ID: data.ID,
		User: &model.UserResponse{
			ID:        data.User.ID,
			Name:      data.User.Name,
			CreatedAt: data.User.CreatedAt,
			UpdatedAt: data.User.UpdatedAt,
		},
		UserID:    data.UserID,
		SourceID:  data.SourceID,
		OreID:     data.OreID,
		Weight:    data.Weight,
		Notes:     data.Notes,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.ProductionResponse]{
		Message: "Production retrived successfully",
		Data:    production,
	}
}

func ProductionsSuccessResponse(data []*entity.Production) *ApiResponse[[]model.ProductionResponse] {
	productions := make([]model.ProductionResponse, len(data))

	for i, production := range data {
		productions[i] = model.ProductionResponse{
			ID:     production.ID,
			UserID: production.UserID,
			User: &model.UserResponse{
				ID:        production.User.ID,
				Name:      production.User.Name,
				CreatedAt: production.User.CreatedAt,
				UpdatedAt: production.User.UpdatedAt,
			},
			SourceID:  production.SourceID,
			OreID:     production.OreID,
			Weight:    production.Weight,
			Notes:     production.Notes,
			CreatedAt: production.CreatedAt,
			UpdatedAt: production.UpdatedAt,
		}
	}

	return &ApiResponse[[]model.ProductionResponse]{
		Message: "Productions retrived successfully",
		Data:    productions,
	}
}

func CreateProductionSuccessResponse(data *entity.Production) *ApiResponse[model.ProductionResponse] {
	production := model.ProductionResponse{
		ID:     data.ID,
		Weight: data.Weight,
		Notes:  data.Notes,
		Ore: &model.OreResponse{
			ID:        data.Ore.ID,
			Name:      data.Ore.Name,
			CreatedAt: data.Ore.CreatedAt,
			UpdatedAt: data.Ore.UpdatedAt,
		},
		Source: &model.SourceResponse{
			ID:        data.Source.ID,
			Name:      data.Source.Name,
			CreatedAt: data.Source.CreatedAt,
			UpdatedAt: data.Source.UpdatedAt,
		},
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.ProductionResponse]{
		Message: "Production created successfully",
		Data:    production,
	}
}

func UpdateProductionSuccessResponse(data *entity.Production) *ApiResponse[model.ProductionResponse] {
	production := model.ProductionResponse{
		ID:       data.ID,
		UserID:   data.UserID,
		SourceID: data.SourceID,
		OreID:    data.OreID,
		Weight:   data.Weight,
		Notes:    data.Notes,
		Ore: &model.OreResponse{
			ID:        data.Ore.ID,
			Name:      data.Ore.Name,
			CreatedAt: data.Ore.CreatedAt,
			UpdatedAt: data.Ore.UpdatedAt,
		},
		Source: &model.SourceResponse{
			ID:        data.Source.ID,
			Name:      data.Source.Name,
			CreatedAt: data.Source.CreatedAt,
			UpdatedAt: data.Source.UpdatedAt,
		},
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.ProductionResponse]{
		Message: "Production updated successfully",
		Data:    production,
	}
}
