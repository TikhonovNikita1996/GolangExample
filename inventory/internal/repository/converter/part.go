package repoConverter

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
	repoModel "github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/repository/model"
)

func ToRepoModelPart(part model.Part) repoModel.Part {
	return repoModel.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      repoModel.Category(part.Category),
		Dimensions:    toRepoDimensions(part.Dimensions),
		Manufacturer:  repoModel.Manufacturer(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      ToRepoMetadata(part.Metadata),
	}
}

func ToModelPart(part repoModel.Part) *model.Part {
	return &model.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      model.Category(part.Category),
		Dimensions:    toModelDimensions(part.Dimensions),
		Manufacturer:  model.Manufacturer(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      ToModelMetadata(part.Metadata),
	}
}

func ToRepoPartsFilter(filter model.PartsFilter) repoModel.PartsFilter {
	return repoModel.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            toRepoCategories(filter.Categories),
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func toRepoDimensions(dimensions model.Dimensions) repoModel.Dimensions {
	return repoModel.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func toModelDimensions(dimensions repoModel.Dimensions) model.Dimensions {
	return model.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func ToRepoMetadata(src map[string]*model.Value) map[string]*repoModel.Value {
	metaData := make(map[string]*repoModel.Value, len(src))
	for k, v := range src {
		metaData[k] = ToRepoValue(v)
	}
	return metaData
}

func toRepoCategories(src []model.Category) []repoModel.Category {
	categories := make([]repoModel.Category, len(src))
	for i, v := range src {
		elem := toRepoCategory(v)
		categories[i] = elem
	}
	return categories
}

func toRepoCategory(category model.Category) repoModel.Category {
	return repoModel.Category(category)
}

func ToModelMetadata(src map[string]*repoModel.Value) map[string]*model.Value {
	metaData := make(map[string]*model.Value, len(src))
	for k, v := range src {
		metaData[k] = toModelValue(v)
	}
	return metaData
}

func ToRepoValue(value *model.Value) *repoModel.Value {
	var stringValue *string
	if value.StringValue != nil {
		tmp := value.StringValue
		stringValue = tmp
	}

	var doubleValue *float64
	if value.DoubleValue != nil {
		tmp := value.DoubleValue
		doubleValue = tmp
	}

	var intValue *int64
	if value.Int64Value != nil {
		tmp := value.Int64Value
		intValue = tmp
	}

	var boolValue *bool
	if value.BoolValue != nil {
		tmp := value.BoolValue
		boolValue = tmp
	}

	return &repoModel.Value{
		StringValue: stringValue,
		DoubleValue: doubleValue,
		BoolValue:   boolValue,
		Int64Value:  intValue,
	}
}

func toModelValue(value *repoModel.Value) *model.Value {
	var stringValue *string
	if value.StringValue != nil {
		tmp := value.StringValue
		stringValue = tmp
	}

	var doubleValue *float64
	if value.DoubleValue != nil {
		tmp := value.DoubleValue
		doubleValue = tmp
	}

	var intValue *int64
	if value.Int64Value != nil {
		tmp := value.Int64Value
		intValue = tmp
	}

	var boolValue *bool
	if value.BoolValue != nil {
		tmp := value.BoolValue
		boolValue = tmp
	}

	return &model.Value{
		StringValue: stringValue,
		DoubleValue: doubleValue,
		BoolValue:   boolValue,
		Int64Value:  intValue,
	}
}
