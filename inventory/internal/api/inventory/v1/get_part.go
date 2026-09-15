package v1

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	converter "github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/converter"
	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
	inventoryV1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
)

func (api *api) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	part, err := api.InventoryService.GetPart(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, model.ErrPartNotFound) {
			return nil, status.Errorf(codes.NotFound, "Part not found")
		}
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}
	return &inventoryV1.GetPartResponse{
		Part: converter.ToProtoPart(*part),
	}, nil
}
