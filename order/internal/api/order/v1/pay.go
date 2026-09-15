package v1

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func (a *OrderAPI) PayForOrderByUuid(ctx context.Context, req *orderv1.PayOrderRequest, params orderv1.PayForOrderByUuidParams) (orderv1.PayForOrderByUuidRes, error) {
	res, err := a.OrderService.PayForOrderByUuid(ctx, req.UserUUID, req.PaymentMethod, params.OrderUUID)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &orderv1.NotFoundError{
				Code:    404,
				Message: "Order with id '" + params.OrderUUID + "' can not found",
			}, nil
			if errors.Is(err, model.ErrOrderCanNotBePayed) {
				return &orderv1.ConflictError{
					Code:    409,
					Message: "Order with id '" + params.OrderUUID + "' can not be payed",
				}, nil
			}
		}
		return &orderv1.InternalServerError{
			Code:    500,
			Message: "Internal Server Error",
		}, nil
	}
	response := &orderv1.PayOrderResponse{
		TransactionUUID: orderv1.NewOptString(res),
	}
	return response, nil
}
