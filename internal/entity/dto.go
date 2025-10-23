package entity

import "time"

type CreateUserRequest struct {
	TelegramID int64   `json:"telegram_id"`
	Phone      *string `json:"phone"`
	Username   *string `json:"username"`
}

type UpdateUserRequest struct {
	Phone    *string `json:"phone"`
	Username *string `json:"username"`
}

type UserResponse struct {
	ID         int64     `json:"id"`
	TelegramID int64     `json:"telegram_id"`
	Phone      *string   `json:"phone"`
	Username   *string   `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateUserResponse struct {
	ID int64 `json:"id"`
}

type ExistUserResponse struct {
	Exists bool `json:"exists"`
}

func ToUserResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:         u.ID,
		TelegramID: u.TelegramID,
		Phone:      u.Phone,
		Username:   u.Username,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
