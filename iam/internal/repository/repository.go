package repository

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/iam/internal/model"
)

type UserRepository interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
	WhoAmI(ctx context.Context, req model.WhoamiRequest) (*model.WhoamiRequest, error)
}

type SessionRepository interface {
	CreateSession(ctx context.Context) (*string, error)
	GetSession(ctx context.Context, id string) (*model.Session, error)
	AddSessionToUseSet(ctx context.Context, userid, sessionId string) error
}
