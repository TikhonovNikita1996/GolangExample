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
	shipAssembledConsumer kafka.Consumer
	shipAssembledDecoder  kafkaConverter.ShipAssembledDecoder
	telegramService       telService.TelegramService
}

func NewService(shipAssembledConsumer kafka.Consumer, shipAssembledDecoder kafkaConverter.ShipAssembledDecoder, telegramService telService.TelegramService) *service {
	return &service{
		shipAssembledConsumer: shipAssembledConsumer,
		shipAssembledDecoder:  shipAssembledDecoder,
		telegramService:       telegramService,
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
