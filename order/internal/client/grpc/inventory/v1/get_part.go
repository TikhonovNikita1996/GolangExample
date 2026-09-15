package v1

import (
	"golang.org/x/net/context"

	clientConverter "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/client/converter"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (c *inventoryServiceClient) GetPart(ctx context.Context, uuid string) (*model.Part, error) {
	request := clientConverter.UuidToProtoPartRequest(uuid)
	part, err := c.generatedClient.GetPart(ctx, request)
	if err != nil {
		return &model.Part{}, err
	}
	return clientConverter.ProtoPartToModel(part.Part), nil
}
