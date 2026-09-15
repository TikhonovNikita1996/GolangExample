package order_consumer

import (
	"context"

	"go.uber.org/zap"

	kafkaConverter "github.com/TikhonovNikita1996/Go-microservises-project/assembly/internal/converter/kafka"
	producerService "github.com/TikhonovNikita1996/Go-microservises-project/assembly/internal/service"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
)

type service struct {
	orderPaidConsumer kafka.Consumer
	orderPaidDecoder  kafkaConverter.OrderPaidDecoder

	shipAssembledProducer producerService.AssemblyProducerService
}

func NewService(orderPaidConsumer kafka.Consumer, orderPaidDecoder kafkaConverter.OrderPaidDecoder, shipAssembledProducer producerService.AssemblyProducerService) *service {
	return &service{
		orderPaidConsumer:     orderPaidConsumer,
		orderPaidDecoder:      orderPaidDecoder,
		shipAssembledProducer: shipAssembledProducer,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order orderPaidConsumer service")

	err := s.orderPaidConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from ufo.recorded topic error", zap.Error(err))
		return err
	}

	return nil
}
