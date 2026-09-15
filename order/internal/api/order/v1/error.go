package v1

import (
	"net/http"

	"golang.org/x/net/context"

	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func (a *OrderAPI) NewError(_ context.Context, err error) *orderv1.GenericErrorStatusCode {
	return &orderv1.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: orderv1.GenericError{
			Code:    orderv1.NewOptInt(http.StatusInternalServerError),
			Message: orderv1.NewOptString(err.Error()),
		},
	}
}
