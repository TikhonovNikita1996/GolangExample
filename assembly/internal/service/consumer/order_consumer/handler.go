package order_consumer

import (
	"math/rand"
	"time"

	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
)

func (s *service) OrderHandler(ctx context.Context, msg kafka.Message) error {
	event, err := s.orderPaidDecoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "Failed to decode OrderPaidMessage", zap.Error(err))
		return err
	}

	logger.Info(ctx, "Processing message",
		zap.String("topic", msg.Topic),
		zap.Any("partition", msg.Partition),
		zap.Any("offset", msg.Offset),
		zap.String("event_uuid", event.EventUuid),
		zap.String("order_uuid", event.OrderUuid),
		zap.String("user_uuid", event.UserUuid),
		zap.String("payment_method", event.PaymentMethod),
		zap.String("transaction_uuid", event.TransactionUuid),
	)

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	s.shipAssembledProducer.ProduceShipAssembled(ctx, event)

	return nil
}
