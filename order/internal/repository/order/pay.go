package repository

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (r *OrderRepository) PayForOrderByUuid(ctx context.Context, request *model.PaymentRequest) error {
	updateModel := model.OrderUpdate{
		OrderUUID:       request.OrderUuid,
		TransactionUUID: &request.TransactionUuid,
		PaymentMethod:   &request.PaymentMethod,
		Status:          &request.PaymentStatus,
	}
	ok := r.UpdateOrder(ctx, updateModel)
	if ok != nil {
		return model.ErrExternal
	}
	return nil
}
