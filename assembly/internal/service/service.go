package service

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/assembly/internal/model"
)

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type AssemblyProducerService interface {
	ProduceShipAssembled(ctx context.Context, event model.OrderPaid) error
}
