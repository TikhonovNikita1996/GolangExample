package service

import (
	"golang.org/x/net/context"
)

type PaymentService interface {
	PayOrder(ctx context.Context, orderUuid, userUuid string) (string, error)
}
