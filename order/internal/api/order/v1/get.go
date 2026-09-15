package v1

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/repository/converter"
	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func (a *OrderAPI) GetOrderByUuid(ctx context.Context, params orderv1.GetOrderByUuidParams) (orderv1.GetOrderByUuidRes, error) {
	order, err := a.OrderService.GetOrderByUuid(ctx, params.OrderUUID)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &orderv1.NotFoundError{
				Code:    404,
				Message: "Order with id '" + params.OrderUUID + "' not found",
			}, nil
		}
		return &orderv1.InternalServerError{
			Code:    500,
			Message: "Internal Server Error",
		}, nil
	}
	return converter.ToProtoOrder(&order), nil
}
