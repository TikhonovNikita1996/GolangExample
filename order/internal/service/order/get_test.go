package order

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (s *ServiceSuite) TestGetOrderSuccess() {
	orderId := gofakeit.UUID()

	order := model.Order{
		OrderUUID:       orderId,
		UserUUID:        gofakeit.UUID(),
		PartsUUID:       []string{gofakeit.UUID(), gofakeit.UUID()},
		TotalPrice:      50,
		TransactionUUID: nil,
		PaymentUUID:     nil,
		Status:          model.Status_WaitingForPayment,
	}

	s.orderRepository.On("GetOrderByUuid", s.ctx, orderId).Return(order, nil)

	res, err := s.service.GetOrderByUuid(s.ctx, orderId)

	s.Require().Equal(orderId, res.OrderUUID)
	s.Require().Nil(err)
}

func (s *ServiceSuite) TestGetOrderNotFound() {
	orderUuid := gofakeit.UUID()

	s.orderRepository.On("GetOrderByUuid", s.ctx, orderUuid).Return(model.Order{}, nil)

	res, err := s.service.GetOrderByUuid(s.ctx, orderUuid)

	s.Require().NotEqual(orderUuid, res.OrderUUID)
	s.Require().Nil(err)
}
