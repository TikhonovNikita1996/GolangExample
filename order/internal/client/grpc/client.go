package grpc

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

type InventoryClient interface {
	GetPart(context.Context, string) (*model.Part, error)
	GetParts(context.Context, model.PartsFilter) ([]*model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, info model.PaymentInfo) (string, error)
}
