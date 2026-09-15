package part

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/repository"
	def "github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/service"
)

var _ def.InventoryService = (*service)(nil)

type service struct {
	inventoryRepository repository.InventoryRepository
}

func NewService(rep repository.InventoryRepository) *service {
	return &service{
		inventoryRepository: rep,
	}
}
