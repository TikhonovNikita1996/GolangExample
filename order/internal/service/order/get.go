package order

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (s *OrderService) GetOrderByUuid(ctx context.Context, uuid string) (model.Order, error) {
	order, err := s.orderRepository.GetOrderByUuid(ctx, uuid)
	if err != nil {
		return model.Order{}, fmt.Errorf("order not found: %w", model.ErrOrderNotFound)
	}
	return order, nil
}
