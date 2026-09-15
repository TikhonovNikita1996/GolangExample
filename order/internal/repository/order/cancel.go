package repository

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (r *OrderRepository) CancelOrderByUuid(ctx context.Context, uuid string) (model.CancelResponse, error) {
	order, err := r.GetOrderByUuid(ctx, uuid)
	if err != nil {
		return model.Cancel_Response_NotFound, err
	}
	if order.Status == model.Status_Paid || order.Status == model.Status_Canceled {
		return model.Cancel_Response_Conflict, nil
	}
	newStatus := model.Status_Canceled
	updateModel := model.OrderUpdate{
		OrderUUID: uuid,
		Status:    &newStatus,
	}
	r.UpdateOrder(ctx, updateModel)
	return model.Cancel_Response_NoContent, nil
}
