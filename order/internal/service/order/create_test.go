package order

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (s *ServiceSuite) TestCreateSuccess() {
	userId := gofakeit.UUID()
	partId := gofakeit.UUID()
	partIds := []string{partId}

	partsFilter := model.PartsFilter{
		Uuids: partIds,
	}

	req := &model.CreateOrderRequest{
		UserUUID:   userId,
		PartUuids:  partIds,
		TotalPrice: 0,
	}

	expectedParts := []*model.Part{
		{
			UUID:          partId,
			Name:          "",
			Description:   "",
			Price:         0,
			StockQuantity: 0,
			Category:      0,
			Dimensions:    model.Dimensions{},
			Manufacturer:  model.Manufacturer{},
			Tags:          nil,
			Metadata:      nil,
		},
	}

	expectedUUID := gofakeit.UUID()

	s.orderRepository.On("CreateOrder", s.ctx, req).Return(expectedUUID, nil)
	s.inventoryMock.On("GetParts", s.ctx, partsFilter).Return(expectedParts, nil)

	res, err := s.service.CreateOrder(s.ctx, req)

	s.Require().Equal(expectedUUID, res.OrderUUID)
	s.Require().Nil(err)
}
