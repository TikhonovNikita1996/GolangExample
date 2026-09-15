package user

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/iam/internal/model"
)

func (r *userRepository) Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error) {
	notifJSON, err := json.Marshal(req.Info.UserInfo.NotificationMethods)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal notification methods: %w", err)
	}

	_, err = r.dbConnection.Exec(ctx,
		`INSERT INTO user_infos (user_uuid, login, email, notification_methods)
         VALUES ($1, $2, $3, $4)`,
		req.Info.Uuid, req.Info.UserInfo.Login, req.Info.UserInfo.Email, notifJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new user to db: %w", err)
	}

	response := &model.RegisterResponse{
		UserUuid: req.Info.Uuid,
	}

	return response, nil
}
