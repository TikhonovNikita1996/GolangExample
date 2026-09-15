package part

import (
	"context"
	"fmt"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
)

func (s *service) GetPart(ctx context.Context, uuid string) (*model.Part, error) {
	part, ok := s.inventoryRepository.GetPart(ctx, uuid)
	if ok != nil {
		return nil, fmt.Errorf("part not found %w", model.ErrPartNotFound)
	}
	return part, nil
}
