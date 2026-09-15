package order_producer

import (
	"context"
	"math/rand"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/TikhonovNikita1996/Go-microservises-project/assembly/internal/model"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
	eventsv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/events/v1"
)

type service struct {
	shipAssembledProducer kafka.Producer
}

func NewService(shipAssembledProducer kafka.Producer) *service {
	return &service{
		shipAssembledProducer: shipAssembledProducer,
	}
}

func (p *service) ProduceShipAssembled(ctx context.Context, event model.OrderPaid) error {
	msg := &eventsv1.ShipAssembled{
		EventUuid:    event.EventUuid,
		OrderUuid:    event.OrderUuid,
		UserUuid:     event.UserUuid,
		BuildTimeSec: int64(rand.Intn(50)),
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal ShipAssembled", zap.Error(err))
		return err
	}

	err = p.shipAssembledProducer.Send(ctx, []byte(event.EventUuid), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish ShipAssembled", zap.Error(err))
		return err
	}

	return nil
}
