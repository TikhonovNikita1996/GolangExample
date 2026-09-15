package part

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
)

func (s *ServiceSuite) TestGetPartByUuidSuccess() {
	uuid := gofakeit.UUID()

	responsePart := model.Part{
		UUID:          uuid,
		Name:          gofakeit.Name(),
		Description:   gofakeit.ProductDescription(),
		Price:         gofakeit.Float64(),
		StockQuantity: gofakeit.Int64(),
		Category:      model.CATEGORY_ENGINE,
		Dimensions:    model.Dimensions{},
		Manufacturer:  model.Manufacturer{},
		Tags:          nil,
		Metadata:      nil,
	}

	s.inventoryRepository.On("GetPart", s.ctx, uuid).Return(responsePart, nil)

	res, err := s.service.GetPart(s.ctx, uuid)

	s.Require().Equal(uuid, res.UUID)
	s.Require().Nil(err)
}

func (s *ServiceSuite) TestGetPartByUuidNotFound() {
	id := gofakeit.UUID()

	s.inventoryRepository.On("GetPart", s.ctx, id).Return(model.Part{}, model.ErrPartNotFound)

	res, err := s.service.GetPart(s.ctx, id)

	s.Require().Equal(model.Part{}, res)
	s.Require().Equal(model.ErrPartNotFound, err)
}
