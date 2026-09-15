package v1

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	converter "github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/converter"
	inventoryV1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
)

func (api *api) GetParts(ctx context.Context, req *inventoryV1.GetPartsRequest) (*inventoryV1.GetPartsResponse, error) {
	parts, err := api.InventoryService.GetParts(ctx, converter.ToModelPartsFilter(req.PartsFilter))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}
	if parts == nil {
		return &inventoryV1.GetPartsResponse{
			Parts: make([]*inventoryV1.Part, 0),
		}, status.Errorf(codes.NotFound, "Part not found")
	}
	return &inventoryV1.GetPartsResponse{
		Parts: converter.ToProtoPartsList(parts),
	}, nil
}
