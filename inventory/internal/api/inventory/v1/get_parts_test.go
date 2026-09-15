package v1

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/converter"
	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
	inventoryv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
)

func (a *ApiSuite) TestGetPartsListSuccess() {
	firstTestPartUuid := gofakeit.UUID()
	secondTestPartUuid := gofakeit.UUID()
	thirdTestPartUuid := gofakeit.UUID()

	responseParts := []model.Part{
		generateTestPart(firstTestPartUuid),
		generateTestPart(secondTestPartUuid),
		generateTestPart(thirdTestPartUuid),
	}
	partsFilter := model.PartsFilter{
		Uuids:      []string{firstTestPartUuid, secondTestPartUuid, thirdTestPartUuid},
		Categories: []model.Category{model.CATEGORY_WING},
	}

	a.service.On("GetParts", a.ctx, partsFilter).Return(responseParts, nil)

	requestFilter := converter.ToProtoPartsFilter(&partsFilter)
	request := &inventoryv1.GetPartsRequest{
		PartsFilter: requestFilter,
	}
	res, err := a.api.GetParts(a.ctx, request)

	a.Require().Equal(converter.ToModelPartsList(res.Parts), responseParts)
	a.Require().Nil(err)
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
