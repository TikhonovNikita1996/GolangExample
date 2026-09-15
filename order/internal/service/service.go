package service

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

type OrderService interface {
	GetOrderByUuid(ctx context.Context, uuid string) (model.Order, error)
	CreateOrder(ctx context.Context, req *model.CreateOrderRequest) (model.CreateOrderResponse, error)
	CancelOrderByUuid(ctx context.Context, uuid string) (model.CancelResponse, error)
	PayForOrderByUuid(ctx context.Context, userUuid, paymentMethod, uuid string) (string, error)
	UpdateOrder(ctx context.Context, update model.OrderUpdate) error
}

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type OrderProducerService interface {
	ProduceOrderPaid(ctx context.Context, event model.OrderPaid) error
}
