package model

type Part struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	Metadata      map[string]*Value
}

type Category int32

const (
	CATEGORY_UNSPECIFIED Category = 0
	CATEGORY_ENGINE      Category = 1
	CATEGORY_FUEL        Category = 2
	CATEGORY_PORTHOLE    Category = 3
	CATEGORY_WING        Category = 4
)

type Value struct {
	StringValue *string
	Int64Value  *int64
	DoubleValue *float64
	BoolValue   *bool
}

type Dimensions struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

type Manufacturer struct {
	Name    string
	Country string
	Website string
}

type PartsFilter struct {
	Uuids                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}
