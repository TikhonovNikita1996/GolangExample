package order_consumer

import (
	"math/rand"
	"time"

	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
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

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	newStatus := model.Status_Assemdled
	updateModel := model.OrderUpdate{
		OrderUUID: event.OrderUuid,
		Status:    &newStatus,
	}

	s.orderService.UpdateOrder(ctx, updateModel)

	return nil
}
