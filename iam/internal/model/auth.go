package model

import (
	"time"

	v1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/common/v1"
)

type LoginRequest struct {
	Login    string
	Password string
}

type LoginResponse struct {
	SessionUuid string
}

type RegisterRequest struct {
	Info     User
	Password string
}

type RegisterResponse struct {
	UserUuid string
}

type WhoamiRequest struct {
	SessionUuid string
}

type WhoamiResponse struct {
	Session *v1.Session
	User    *v1.User
}

type Session struct {
	Uuid      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	ExpiresAt *time.Time
}
