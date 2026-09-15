package converter

import (
	"github.com/samber/lo"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	inventoryv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
)

func ProtoPartToModel(part *inventoryv1.Part) *model.Part {
	return &model.Part{
		UUID:          part.GetUuid(),
		Name:          part.GetName(),
		Description:   part.GetDescription(),
		Price:         part.Price,
		StockQuantity: part.GetStockQuantity(),
		Category:      model.Category(part.Category),
		Dimensions:    toModelProtoDimensions(part.Dimensions),
		Manufacturer:  model.Manufacturer{},
		Tags:          part.GetTags(),
		Metadata:      toModelProtoMetaData(part.Metadata),
	}
}

func PartToProto(part model.Part) *inventoryv1.Part {
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

func ProtoPartsFilterToModel(filter *inventoryv1.PartsFilter) model.PartsFilter {
	return model.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            toModelProtoCategories(filter.Categories),
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func PartsFilterToProto(filter model.PartsFilter) *inventoryv1.PartsFilter {
	return &inventoryv1.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            toProtoCategories(filter.Categories),
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func ListOfPartsToProto(parts []model.Part) []*inventoryv1.Part {
	result := make([]*inventoryv1.Part, 0, len(parts))
	for i, p := range parts {
		result[i] = PartToProto(p)
	}
	return result
}

func ToModelListOfParts(parts []*inventoryv1.Part) []*model.Part {
	result := make([]*model.Part, 0, len(parts))
	for _, p := range parts {
		result = append(result, ProtoPartToModel(p))
	}
	return result
}

func UuidToProtoPartRequest(uuid string) *inventoryv1.GetPartRequest {
	result := inventoryv1.GetPartRequest{
		Uuid: uuid,
	}
	return &result
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

func toModelProtoDimensions(dimensions *inventoryv1.Dimensions) model.Dimensions {
	return model.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func toModelProtoMetaData(src map[string]*inventoryv1.Value) map[string]*model.Value {
	dst := make(map[string]*model.Value, len(src))
	for k, v := range src {
		dst[k] = toModelProtoValue(v)
	}
	return dst
}

func toProtoMetadata(src map[string]*model.Value) map[string]*inventoryv1.Value {
	dst := make(map[string]*inventoryv1.Value, len(src))
	for k, v := range src {
		dst[k] = toProtoValue(v)
	}
	return dst
}

func toModelProtoValue(value *inventoryv1.Value) *model.Value {
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
	default:
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

func toModelProtoCategory(category inventoryv1.Category) model.Category {
	return model.Category(category)
}

func toProtoCategory(category model.Category) inventoryv1.Category {
	return inventoryv1.Category(category)
}

func toModelProtoCategories(src []inventoryv1.Category) []model.Category {
	dst := make([]model.Category, len(src))
	for i, v := range src {
		elem := toModelProtoCategory(v)
		dst[i] = elem
	}
	return dst
}

func toProtoCategories(src []model.Category) []inventoryv1.Category {
	dst := make([]inventoryv1.Category, len(src))
	for i, v := range src {
		elem := toProtoCategory(v)
		dst[i] = elem
	}
	return dst
}
