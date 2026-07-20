package auth

import "portfolio-backend/internal/modules/admin"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	Admin admin.Admin `json:"admin"`
}
