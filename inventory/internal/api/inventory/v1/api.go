package v1

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/service"
	inventoryv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
)

type api struct {
	inventoryv1.UnimplementedInventoryServiceServer

	InventoryService service.InventoryService
}

func NewAPI(inventoryService service.InventoryService) *api {
	return &api{
		InventoryService: inventoryService,
	}
}
