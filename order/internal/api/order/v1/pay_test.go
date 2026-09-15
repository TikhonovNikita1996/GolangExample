package v1

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func (a *ApiSuite) TestPaymentSuccess() {
	id := gofakeit.UUID()
	transactionUuid := gofakeit.UUID()
	updateModel := model.OrderUpdate{
		TransactionUUID: &transactionUuid,
	}

	a.service.On("GetOrderByUuid", a.ctx, id).Return(model.Order{}, nil)
	a.service.On("UpdateOrder", a.ctx, id, updateModel).Return(nil)
	a.service.On("PayForOrderByUuid", a.ctx, "CARD", id).Return(transactionUuid, nil)

	payParams := orderv1.PayForOrderByUuidParams{OrderUUID: id}
	payRequest := &orderv1.PayOrderRequest{
		PaymentMethod: "CARD",
	}

	expectedResponse := orderv1.PayOrderResponse{TransactionUUID: orderv1.NewOptString(transactionUuid)}
	res, err := a.api.PayForOrderByUuid(a.ctx, payRequest, payParams)

	a.Require().Equal(res, &expectedResponse)
	a.Require().Nil(err)
}
