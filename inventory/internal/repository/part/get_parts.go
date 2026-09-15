package part

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/helpers"
	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
	repoConverter "github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/repository/converter"
	repoModel "github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/repository/model"
)

func (r *repository) GetParts(ctx context.Context, filter model.PartsFilter) ([]*model.Part, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, model.ErrInternal
	}
	defer func() {
		cerr := cursor.Close(ctx)
		if cerr != nil {
			log.Printf("failed to close cursor: %v\n", cerr)
		}
	}()

	var parts []repoModel.Part
	err = cursor.All(ctx, &parts)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Part, 0, len(parts))
	filterParts := helpers.FilterParts(parts, repoConverter.ToRepoPartsFilter(filter))
	for _, part := range filterParts {
		result = append(result, repoConverter.ToModelPart(part))
	}

	return result, nil
}
