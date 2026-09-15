package v1

import (
	"golang.org/x/net/context"

	clientConverter "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/client/converter"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	generatedInventoryv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
)

func (c *inventoryServiceClient) GetParts(ctx context.Context, filter model.PartsFilter) ([]*model.Part, error) {
	parts, err := c.generatedClient.GetParts(ctx, &generatedInventoryv1.GetPartsRequest{
		PartsFilter: clientConverter.PartsFilterToProto(filter),
	})
	if err != nil {
		return nil, err
	}
	return clientConverter.ToModelListOfParts(parts.Parts), nil
}
