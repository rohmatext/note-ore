package presenter

import (
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
)

func UserSuccessResponse(data *entity.User) *ApiResponse[model.UserResponse] {
	user := model.UserResponse{
		ID:        data.ID,
		Name:      data.Name,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &ApiResponse[model.UserResponse]{
		Message: "User data retrived successfully",
		Data:    user,
	}
}

func UsersSuccessResponse(data []*entity.User) *ApiResponse[[]model.UserResponse] {
	users := make([]model.UserResponse, len(data))

	for i, user := range data {
		users[i] = model.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return &ApiResponse[[]model.UserResponse]{
		Message: "User data retrived successfully",
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
