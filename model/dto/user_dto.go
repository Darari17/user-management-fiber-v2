package dto

import "time"

// DTO ini digunakan untuk request pembuatan user baru
type CreateRequest struct {
	Username string `validate:"required,min=5,max=255"`
	Email    string `validate:"required,email,max=255,min=5"`
	Password string `validate:"required,min=5,max=255"`
}

// DTO ini digunakan untuk mengupdate data user.
type UpdateRequest struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"omitempty,min=5,max=255"`
	Email    string `json:"email" validate:"omitempty,email,max=255,min=5"`
	Password string `json:"password" validate:"omitempty,min=5,max=255"`
}

// DTO ini digunakan untuk mengembalikan data user sebagai response API
type Response struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// DTO ini digunakan sebagai struktur standar response API.
type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}
