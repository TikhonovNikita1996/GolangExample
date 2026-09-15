package part

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
)

func (s *ServiceSuite) TestGetPartsListSuccess() {
	firstTestPartUuid := gofakeit.UUID()
	secondTestPartUuid := gofakeit.UUID()
	thirdTestPartUuid := gofakeit.UUID()

	responseParts := []model.Part{
		generateTestPart(firstTestPartUuid),
		generateTestPart(secondTestPartUuid),
		generateTestPart(thirdTestPartUuid),
	}
	partsFilter := model.PartsFilter{
		Uuids: []string{firstTestPartUuid, secondTestPartUuid, thirdTestPartUuid},
	}

	s.inventoryRepository.On("GetParts", s.ctx, partsFilter).Return(responseParts, nil)

	res, err := s.service.GetParts(s.ctx, partsFilter)

	s.Require().Equal(responseParts, res)
	s.Require().Nil(err)
}

func generateTestPart(uuid string) model.Part {
	return model.Part{
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
}
