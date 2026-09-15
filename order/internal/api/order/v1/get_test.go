package v1

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func (a *ApiSuite) TestGetOrderSuccess() {
	id := gofakeit.UUID()

	request := orderv1.GetOrderByUuidParams{OrderUUID: id}
	order := model.Order{
		OrderUUID:       id,
		UserUUID:        gofakeit.UUID(),
		PartsUUID:       []string{gofakeit.UUID(), gofakeit.UUID()},
		TotalPrice:      50,
		TransactionUUID: &id,
		PaymentUUID:     &id,
		Status:          model.Status_WaitingForPayment,
	}

	a.service.On("GetOrderByUuid", a.ctx, id).Return(order, nil)

	res, err := a.api.GetOrderByUuid(a.ctx, request)

	resOrder, _ := res.(*orderv1.Order)

	a.Require().Equal(id, resOrder.OrderUUID)
	a.Require().Nil(err)
}
