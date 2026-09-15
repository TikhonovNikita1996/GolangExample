package service

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/model"
)

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type TelegramService interface {
	SendPaidNotification(ctx context.Context, model model.OrderPaid) error
	SendShipAssembledNotification(ctx context.Context, model model.ShipAssembled) error
}
