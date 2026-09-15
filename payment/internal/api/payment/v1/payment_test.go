package v1

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/payment/internal/model"
	paymentv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/payment/v1"
)

func (a *ApiSuite) TestPaySuccess() {
	orderUuid := gofakeit.UUID()
	userUuid := gofakeit.UUID()

	req := &paymentv1.PayOrderRequest{
		Info: &paymentv1.PaymentInfo{
			OrderUuid:     orderUuid,
			UserUuid:      userUuid,
			PaymentMethod: 1,
		},
	}

	a.service.On("PayOrder", a.ctx, orderUuid, userUuid).Return("", nil)

	res, err := a.api.PayOrder(a.ctx, req)
	a.Require().NoError(err)
	a.Require().NotNil(res)
}

func (a *ApiSuite) TestPayWithoutUserUuidError() {
	orderUuid := gofakeit.UUID()

	req := &paymentv1.PayOrderRequest{
		Info: &paymentv1.PaymentInfo{
			OrderUuid:     orderUuid,
			UserUuid:      "",
			PaymentMethod: 1,
		},
	}

	a.service.On("PayOrder", a.ctx, orderUuid, "").Return("", model.ErrArgumentEmptyString)

	res, err := a.api.PayOrder(a.ctx, req)
	a.Require().Error(err)
	a.Require().Nil(res)
}
