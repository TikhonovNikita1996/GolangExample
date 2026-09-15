package v1

import (
	inventoryv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
)

type inventoryServiceClient struct {
	generatedClient inventoryv1.InventoryServiceClient
}

func NewClient(generatedClient inventoryv1.InventoryServiceClient) *inventoryServiceClient {
	return &inventoryServiceClient{generatedClient}
}
