package converter

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/iam/internal/model"
	repoModel "github.com/TikhonovNikita1996/Go-microservises-project/iam/internal/repository/model"
	"github.com/samber/lo"
)

func ToModelViewSession(repoModel repoModel.SessionRedisView) (model.Session, error) {
	var createdAtNs *int64
	if repoModel.CreatedAt != nil {
		createdAtNs = lo.ToPtr(repoModel.CreatedAt.UnixNano())
	}

	var updatedAtNs *int64
	if repoModel.UpdatedAt != nil {
		updatedAtNs = lo.ToPtr(repoModel.UpdatedAt.UnixNano())
	}

	var expiresAt *int64
	if repoModel.ExpiresAt != nil {
		expiresAt = lo.ToPtr(repoModel.ExpiresAt.UnixNano())
	}

	return model.Session{
		Uuid:      "",
		CreatedAt: nil,
		UpdatedAt: nil,
		ExpiresAt: nil,
	}, nil
}


func ToRedisViewSession(model model.Session) (repoModel.SessionRedisView, error) {
	var createdAtNs *int64
	if model.CreatedAt != nil {
		createdAtNs = lo.ToPtr(model.CreatedAt.UnixNano())
	}

	var updatedAtNs *int64
	if model.UpdatedAt != nil {
		updatedAtNs = lo.ToPtr(model.UpdatedAt.UnixNano())
	}

	var expiresAt *int64
	if model.ExpiresAt != nil {
		expiresAt = lo.ToPtr(model.ExpiresAt.UnixNano())
	}

	return repoModel.SessionRedisView{
		UUID:        "",
		CreatedAtNs: createdAtNs,
		UpdatedAtNs: updatedAtNs,
		ExpiresAt:   expiresAt,
	}, nil
}
