package v1

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
	inventoryv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
)

func (a *ApiSuite) TestGetPartByUuidSuccess() {
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

	a.service.On("GetPart", a.ctx, uuid).Return(responsePart, nil)

	request := &inventoryv1.GetPartRequest{
		Uuid: uuid,
	}

	res, err := a.api.GetPart(a.ctx, request)

	a.Require().Equal(uuid, res.Part.Uuid)
	a.Require().Nil(err)
}

func (a *ApiSuite) TestGetPartByUuidNotFound() {
	uuid := gofakeit.UUID()
	request := &inventoryv1.GetPartRequest{
		Uuid: uuid,
	}

	a.service.On("GetPart", a.ctx, uuid).Return(model.Part{}, model.ErrPartNotFound)
	_, err := a.api.GetPart(a.ctx, request)

	a.Require().Equal(err, model.ErrPartNotFound)
}
