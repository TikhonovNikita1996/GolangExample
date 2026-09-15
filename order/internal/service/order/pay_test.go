package order

import (
	"github.com/brianvoe/gofakeit/v7"
)

func (s *ServiceSuite) TestPaymentSuccess() {
	id := gofakeit.UUID()
	s.orderRepository.On("PayForOrderByUuid", s.ctx, "CARD", id).Return("", nil)

	_, err := s.service.PayForOrderByUuid(s.ctx, "CARD", id)

	s.Require().Nil(err)
}
