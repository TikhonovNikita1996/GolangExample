package repository

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

type OrderRepository interface {
	GetOrderByUuid(ctx context.Context, uuid string) (model.Order, error)
	CreateOrder(ctx context.Context, req *model.CreateOrderRequest) (string, error)
	CancelOrderByUuid(ctx context.Context, uuid string) (model.CancelResponse, error)
	PayForOrderByUuid(ctx context.Context, request *model.PaymentRequest) error
	UpdateOrder(ctx context.Context, update model.OrderUpdate) error
}
