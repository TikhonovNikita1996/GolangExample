package producer

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
	eventsv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/events/v1"
)

type service struct {
	orderPaidProducer kafka.Producer
}

func NewService(orderPaidProducer kafka.Producer) *service {
	return &service{
		orderPaidProducer: orderPaidProducer,
	}
}

func (p *service) ProduceOrderPaid(ctx context.Context, event model.OrderPaid) error {
	msg := &eventsv1.OrderPayed{
		EventUuid:       event.EventUuid,
		OrderUuid:       event.OrderUuid,
		UserUuid:        event.UserUuid,
		PaymentMethod:   event.PaymentMethod,
		TransactionUuid: event.TransactionUuid,
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal OrderPaid", zap.Error(err))
		return err
	}

	err = p.orderPaidProducer.Send(ctx, []byte(event.EventUuid), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish OrderPaid", zap.Error(err))
		return err
	}

	return nil
}
