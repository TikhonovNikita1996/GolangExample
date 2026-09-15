package v1

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/converter"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func (a *OrderAPI) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (orderv1.CreateOrderRes, error) {
	creationResponse, err := a.OrderService.CreateOrder(ctx, converter.ToModelOrderCreateRequest(req))
	if err != nil {
		if errors.Is(err, model.ErrPartsNotAvailable) {
			return &orderv1.NotFoundError{
				Code:    404,
				Message: "Parts Not Available",
			}, nil
		}
		return &orderv1.InternalServerError{
			Code:    500,
			Message: "Internal Server Error",
		}, nil
	}
	return converter.ToProtoOrderCreateResponse(creationResponse), nil
}
