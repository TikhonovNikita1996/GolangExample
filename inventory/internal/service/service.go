package service

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
)

type InventoryService interface {
	GetPart(context.Context, string) (*model.Part, error)
	GetParts(context.Context, model.PartsFilter) ([]*model.Part, error)
}
