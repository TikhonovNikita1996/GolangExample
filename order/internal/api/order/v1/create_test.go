package v1

import (
	"github.com/brianvoe/gofakeit/v7"

	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func (a *ApiSuite) TestCreateSuccess() {
	orderUuid := gofakeit.UUID()
	parstUuids := []string{gofakeit.UUID(), gofakeit.UUID()}

	request := &orderv1.CreateOrderRequest{
		UserUUID:  gofakeit.UUID(),
		PartUuids: parstUuids,
	}

	expectedResponse := orderv1.CreateOrderResponse{
		OrderUUID:  orderUuid,
		TotalPrice: 0,
	}

	a.service.On("CreateOrder", a.ctx, orderUuid).Return(expectedResponse, nil)
	res, err := a.api.CreateOrder(a.ctx, request)
	resOrder, _ := res.(*orderv1.CreateOrderResponse)

	a.Require().Equal(resOrder.OrderUUID, &expectedResponse.OrderUUID)
	a.Require().Nil(err)
}
