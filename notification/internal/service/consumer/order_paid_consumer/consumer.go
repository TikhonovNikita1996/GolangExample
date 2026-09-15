package order_consumer

import (
	"context"

	"go.uber.org/zap"

	kafkaConverter "github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/converter/kafka"
	telService "github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/service"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
)

type service struct {
	orderPaidConsumer kafka.Consumer
	orderPaidDecoder  kafkaConverter.OrderPaidDecoder
	telegramService   telService.TelegramService
}

func NewService(orderPaidConsumer kafka.Consumer, orderPaidDecoder kafkaConverter.OrderPaidDecoder, telegramService telService.TelegramService) *service {
	return &service{
		orderPaidConsumer: orderPaidConsumer,
		orderPaidDecoder:  orderPaidDecoder,
		telegramService:   telegramService,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order orderPaidConsumer service")

	err := s.orderPaidConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.paid topic error", zap.Error(err))
		return err
	}

	return nil
}
