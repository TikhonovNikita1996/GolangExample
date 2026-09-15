package user

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/iam/internal/model"
)

func (r *userRepository) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	login := req.Login
	password := req.Password

	row := r.dbConnection.QueryRow(ctx, `
        SELECT u.login, u.password
        FROM users u
        WHERE u.login = $1 AND u.password = $2
    `, req.Login, req.Password)

	err := row.Scan(&login, &password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to scan user info: %w", err)
	}

	response := &model.LoginResponse{
		SessionUuid: "",
	}
	return response, nil
}
