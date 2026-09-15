package order

import (
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
)

func (s *OrderService) CancelOrderByUuid(ctx context.Context, id string) (model.CancelResponse, error) {
	cancelResponse, err := s.orderRepository.CancelOrderByUuid(ctx, id)
	if err != nil {
		logger.Error(ctx, "failed to cancel order", zap.String("id", id), zap.Error(err))
		return model.Cancel_Response_NotFound, fmt.Errorf("cancel order by id %s not found", id)
	}
	return cancelResponse, nil
}
