package order

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (s *ServiceSuite) TestCancelSuccess() {
	id := gofakeit.UUID()

	s.orderRepository.On("CancelOrderByUuid", s.ctx, id).Return(model.Cancel_Response_NoContent, nil)

	res, err := s.service.CancelOrderByUuid(s.ctx, id)

	s.Require().Equal(model.Cancel_Response_NoContent, res)
	s.Require().Nil(err)
}

func (s *ServiceSuite) TestCancelNotFoundError() {
	id := gofakeit.UUID()

	s.orderRepository.On("CancelOrderByUuid", s.ctx, id).Return(model.Cancel_Response_NotFound, nil)

	res, err := s.service.CancelOrderByUuid(s.ctx, id)

	s.Require().Equal(model.Cancel_Response_NotFound, res)
	s.Require().Nil(err)
}

func (s *ServiceSuite) TestCancelAlreadyPayedError() {
	id := gofakeit.UUID()

	s.orderRepository.On("CancelOrderByUuid", s.ctx, id).Return(model.Cancel_Response_Conflict, nil)

	res, err := s.service.CancelOrderByUuid(s.ctx, id)

	s.Require().Equal(model.Cancel_Response_Conflict, res)
	s.Require().Nil(err)
}
