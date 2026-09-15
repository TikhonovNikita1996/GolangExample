package v1

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func (a *ApiSuite) TestCancelSuccess() {
	id := gofakeit.UUID()

	request := orderv1.CancelOrderByUuidParams{OrderUUID: id}

	a.service.On("CancelOrderByUuid", a.ctx, id).Return(model.Cancel_Response_NoContent, nil)

	res, err := a.api.CancelOrderByUuid(a.ctx, request)

	a.Require().Equal(&orderv1.NoContentResponse{
		Code:    204,
		Message: "Order with id '" + id + "' was cancelled",
	}, res)
	a.Require().Nil(err)
}

func (a *ApiSuite) TestCancelNotFound() {
	id := gofakeit.UUID()

	request := orderv1.CancelOrderByUuidParams{OrderUUID: id}

	a.service.On("CancelOrderByUuid", a.ctx, id).Return(model.Cancel_Response_NotFound, nil)

	res, err := a.api.CancelOrderByUuid(a.ctx, request)

	a.Require().Equal(&orderv1.NotFoundError{
		Code:    404,
		Message: "Order with id '" + id + "' can not be found",
	}, res)
	a.Require().Nil(err)
}

func (a *ApiSuite) TestCannotCancel() {
	id := gofakeit.UUID()

	request := orderv1.CancelOrderByUuidParams{OrderUUID: id}

	a.service.On("CancelOrderByUuid", a.ctx, id).Return(model.Cancel_Response_Conflict, nil)

	res, err := a.api.CancelOrderByUuid(a.ctx, request)

	a.Require().Equal(&orderv1.ConflictError{
		Code:    409,
		Message: "Order with id '" + id + "' can not be cancelled",
	}, res)
	a.Require().Nil(err)
}
