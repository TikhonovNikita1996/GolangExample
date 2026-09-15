package v1

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func (a *OrderAPI) CancelOrderByUuid(ctx context.Context, params orderv1.CancelOrderByUuidParams) (orderv1.CancelOrderByUuidRes, error) {
	cancelResponse, _ := a.OrderService.CancelOrderByUuid(ctx, params.OrderUUID)

	switch cancelResponse {
	case model.Cancel_Response_NotFound:
		return &orderv1.NotFoundError{
			Code:    404,
			Message: "Order with id '" + params.OrderUUID + "' can not be found",
		}, nil
	case model.Cancel_Response_Conflict:
		return &orderv1.ConflictError{
			Code:    409,
			Message: "Order with id '" + params.OrderUUID + "' can not be cancelled",
		}, nil
	default:
		return &orderv1.NoContentResponse{
			Code:    204,
			Message: "Order with id '" + params.OrderUUID + "' was cancelled",
		}, nil
	}
}
