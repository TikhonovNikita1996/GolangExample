package order

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (s *OrderService) UpdateOrder(ctx context.Context, updateModel model.OrderUpdate) error {
	updateErr := s.orderRepository.UpdateOrder(ctx, updateModel)
	if updateErr != nil {
		return fmt.Errorf("order not found %w", model.ErrOrderNotFound)
	}
	return nil
}
