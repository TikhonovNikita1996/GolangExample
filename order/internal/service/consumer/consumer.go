package order_consumer

import (
	"context"

	"go.uber.org/zap"

	kafkaConverter "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/converter/kafka"
	orderService "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/service"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
)

type service struct {
	shipAssembledConsumer kafka.Consumer
	shipAssembledDecoder  kafkaConverter.ShipAssembledDecoder
	orderService          orderService.OrderService
}

func NewService(shipAssembledConsumer kafka.Consumer, shipAssembledDecoder kafkaConverter.ShipAssembledDecoder, orderService orderService.OrderService) *service {
	return &service{
		shipAssembledConsumer: shipAssembledConsumer,
		shipAssembledDecoder:  shipAssembledDecoder,
		orderService:          orderService,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order shipAssembledConsumer service")

	err := s.shipAssembledConsumer.Consume(ctx, s.ShipAssembledHandler)
	if err != nil {
		logger.Error(ctx, "Consume from ufo.recorded topic error", zap.Error(err))
		return err
	}

	return nil
}
