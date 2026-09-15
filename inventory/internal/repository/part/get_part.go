package part

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
	repoConverter "github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/repository/converter"
	repoModel "github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/repository/model"
)

func (r *repository) GetPart(ctx context.Context, id string) (*model.Part, error) {
	var part repoModel.Part
	err := r.collection.FindOne(ctx, bson.M{"uuid": id}).Decode(&part)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &model.Part{}, model.ErrPartNotFound
		}
		return &model.Part{}, model.ErrInternal
	}
	return repoConverter.ToModelPart(part), nil
}
