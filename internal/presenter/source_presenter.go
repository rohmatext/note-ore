package presenter

import (
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
)

func SourceSuccessResponse(data *entity.Source) *ApiResponse[model.SourceResponse] {
	source := model.SourceResponse{
		ID:          data.ID,
		Name:        data.Name,
		PhoneNumber: data.PhoneNumber,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return &ApiResponse[model.SourceResponse]{
		Message: "Source retrived successfully",
		Data:    source,
	}
}

func SourcesSuccessResponse(data []*entity.Source) *ApiResponse[[]model.SourceResponse] {
	sources := make([]model.SourceResponse, len(data))

	for i, source := range data {
		sources[i] = model.SourceResponse{
			ID:          source.ID,
			Name:        source.Name,
			PhoneNumber: source.PhoneNumber,
			CreatedAt:   source.CreatedAt,
			UpdatedAt:   source.UpdatedAt,
		}
	}

	return &ApiResponse[[]model.SourceResponse]{
		Message: "Sources retrived successfully",
		Data:    sources,
	}
}

func CreateSourceSuccessResponse(data *entity.Source) *ApiResponse[model.SourceResponse] {
	source := model.SourceResponse{
		ID:          data.ID,
		Name:        data.Name,
		PhoneNumber: data.PhoneNumber,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return &ApiResponse[model.SourceResponse]{
		Message: "Source created successfully",
		Data:    source,
	}
}

func UpdateSourceSuccessResponse(data *entity.Source) *ApiResponse[model.SourceResponse] {
	source := model.SourceResponse{
		ID:          data.ID,
		Name:        data.Name,
		PhoneNumber: data.PhoneNumber,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return &ApiResponse[model.SourceResponse]{
		Message: "Source updated successfully",
		Data:    source,
	}
}
