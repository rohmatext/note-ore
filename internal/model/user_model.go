package model

import (
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	ID        uint          `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	Username  string        `json:"username,omitempty"`
	Role      *RoleResponse `json:"role,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=50"`
	Password string `json:"password" validate:"required"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Username string `json:"username" validate:"required,unique_table=users.username,max=50"`
	Password string `json:"password" validate:"required,min=5"`
}

type UpdateUserRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required,max=100"`
	Username string `json:"username" validate:"required,unique_table=users.username,max=50"`
}

func UserCursorPaginate(limit int, cursorID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Limit(limit + 1).Order("id DESC")
		if cursorID != 0 {
			db = db.Where("id <= ?", cursorID)
		}

		return db
	}
}
