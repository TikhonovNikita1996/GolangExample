package order_consumer

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"

	modelSA "github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/model"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
)

func (s *service) ShipAssembledHandler(ctx context.Context, msg kafka.Message) error {
	event, err := s.shipAssembledDecoder.Decode(msg.Value)
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
		zap.Int64("build_time_sec", event.BuildTimeSec),
	)

	model := modelSA.ShipAssembled{
		EventUuid:    event.EventUuid,
		OrderUuid:    event.OrderUuid,
		UserUuid:     event.UserUuid,
		BuildTimeSec: event.BuildTimeSec,
	}

	s.telegramService.SendShipAssembledNotification(ctx, model)

	return nil
}
