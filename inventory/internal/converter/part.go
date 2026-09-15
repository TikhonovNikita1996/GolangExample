package converter

import (
	"github.com/samber/lo"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
	inventoryv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
)

func ToModelParts(part *inventoryv1.Part) model.Part {
	return model.Part{
		UUID:          part.GetUuid(),
		Name:          part.GetName(),
		Description:   part.GetDescription(),
		Price:         part.Price,
		StockQuantity: part.GetStockQuantity(),
		Category:      model.Category(part.Category),
		Dimensions:    toModelDimensions(part.Dimensions),
		Manufacturer:  model.Manufacturer{},
		Tags:          part.GetTags(),
		Metadata:      toModelMetadata(part.Metadata),
	}
}

func ToProtoPart(part model.Part) *inventoryv1.Part {
	return &inventoryv1.Part{
		Uuid:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      inventoryv1.Category(part.Category),
		Dimensions:    toProtoDimensions(part.Dimensions),
		Manufacturer:  toProtoManufacturer(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      toProtoMetadata(part.Metadata),
	}
}

func ToModelPartsFilter(filter *inventoryv1.PartsFilter) model.PartsFilter {
	return model.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            toModelCategories(filter.Categories),
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func ToProtoPartsFilter(filter *model.PartsFilter) *inventoryv1.PartsFilter {
	return &inventoryv1.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            toProtoCategories(filter.Categories),
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func ToProtoPartsList(parts []*model.Part) []*inventoryv1.Part {
	result := make([]*inventoryv1.Part, 0, len(parts))
	for _, p := range parts {
		result = append(result, ToProtoPart(*p))
	}
	return result
}

func ToModelPartsList(parts []*inventoryv1.Part) []model.Part {
	result := make([]model.Part, 0, len(parts))
	for i, p := range parts {
		result[i] = ToModelParts(p)
	}
	return result
}

func toProtoManufacturer(m model.Manufacturer) *inventoryv1.Manufacturer {
	return &inventoryv1.Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func toProtoDimensions(dimensions model.Dimensions) *inventoryv1.Dimensions {
	return &inventoryv1.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func toModelDimensions(dimensions *inventoryv1.Dimensions) model.Dimensions {
	return model.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func toModelMetadata(src map[string]*inventoryv1.Value) map[string]*model.Value {
	metadata := make(map[string]*model.Value, len(src))
	for k, v := range src {
		metadata[k] = toValueModel(v)
	}
	return metadata
}

func toProtoMetadata(src map[string]*model.Value) map[string]*inventoryv1.Value {
	metaData := make(map[string]*inventoryv1.Value, len(src))
	for k, v := range src {
		metaData[k] = toProtoValue(v)
	}
	return metaData
}

func toValueModel(value *inventoryv1.Value) *model.Value {
	var cv model.Value

	switch value.Kind.(type) {
	case *inventoryv1.Value_StringValue:
		cv.StringValue = lo.ToPtr(value.GetStringValue())
	case *inventoryv1.Value_Int64Value:
		cv.Int64Value = lo.ToPtr(value.GetInt64Value())
	case *inventoryv1.Value_DoubleValue:
		cv.DoubleValue = lo.ToPtr(value.GetDoubleValue())
	case *inventoryv1.Value_BoolValue:
		cv.BoolValue = lo.ToPtr(value.GetBoolValue())
	}
	return &cv
}

func toProtoValue(v *model.Value) *inventoryv1.Value {
	if v == nil {
		return nil
	}

	pv := &inventoryv1.Value{}

	switch {
	case v.StringValue != nil:
		pv.Kind = &inventoryv1.Value_StringValue{StringValue: *v.StringValue}
	case v.Int64Value != nil:
		pv.Kind = &inventoryv1.Value_Int64Value{Int64Value: *v.Int64Value}
	case v.DoubleValue != nil:
		pv.Kind = &inventoryv1.Value_DoubleValue{DoubleValue: *v.DoubleValue}
	case v.BoolValue != nil:
		pv.Kind = &inventoryv1.Value_BoolValue{BoolValue: *v.BoolValue}
	default:
	}

	return pv
}

func toModelCategory(category inventoryv1.Category) model.Category {
	return model.Category(category)
}

func toModelCategories(src []inventoryv1.Category) []model.Category {
	categories := make([]model.Category, len(src))
	for i, v := range src {
		elem := toModelCategory(v)
		categories[i] = elem
	}
	return categories
}

func toProtoCategory(category model.Category) inventoryv1.Category {
	return inventoryv1.Category(category)
}

func toProtoCategories(src []model.Category) []inventoryv1.Category {
	categories := make([]inventoryv1.Category, len(src))
	for i, v := range src {
		elem := toProtoCategory(v)
		categories[i] = elem
	}
	return categories
}
