package presenter

import (
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
)

func UserSuccessResponse(data *entity.User) *ApiResponse[model.UserResponse] {
	user := model.UserResponse{
		ID:       data.ID,
		Name:     data.Name,
		Username: data.Username,
		Role: &model.RoleResponse{
			ID:   data.Role.ID,
			Name: data.Role.Name,
		},
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.UserResponse]{
		Message: "User retrived successfully",
		Data:    user,
	}
}

func CreateUserSuccessResponse(data *entity.User) *ApiResponse[model.UserResponse] {
	user := model.UserResponse{
		ID:        data.ID,
		Name:      data.Name,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.UserResponse]{
		Message: "User created successfully",
		Data:    user,
	}
}

func UpdateUserSuccessResponse(data *entity.User) *ApiResponse[model.UserResponse] {
	user := model.UserResponse{
		ID:        data.ID,
		Name:      data.Name,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.UserResponse]{
		Message: "User updated successfully",
		Data:    user,
	}
}

func UsersSuccessResponse(data []*entity.User) *ApiResponse[[]model.UserResponse] {
	users := make([]model.UserResponse, len(data))

	for i, user := range data {
		users[i] = model.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Role: &model.RoleResponse{
				ID:   user.Role.ID,
				Name: user.Role.Name,
			},
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return &ApiResponse[[]model.UserResponse]{
		Message: "Users retrived successfully",
		Data:    users,
	}
}

func UserLoginSuccessResponse(tokenString string) *ApiResponse[model.TokenResponse] {
	token := model.TokenResponse{
		Token: tokenString,
	}

	return &ApiResponse[model.TokenResponse]{
		Message: "User logged in successfully",
		Data:    token,
	}
}
