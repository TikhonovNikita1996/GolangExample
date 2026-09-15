package part

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
)

func (s *service) GetParts(ctx context.Context, filter model.PartsFilter) ([]*model.Part, error) {
	parts, err := s.inventoryRepository.GetParts(ctx, filter)
	if err != nil {
		return []*model.Part{}, fmt.Errorf("parts not available: %w", err)
	}
	return parts, nil
}
