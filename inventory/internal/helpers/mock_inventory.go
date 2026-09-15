package helpers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/repository/model"
)

func SeedParts(ctx context.Context, collection *mongo.Collection) error {
	parts := []interface{}{
		model.Part{
			UUID:          "a5698d66-eba0-4805-9288-5e846340bc49",
			Name:          "Ion Engine X1",
			Description:   "High-efficiency ion propulsion system.",
			Price:         120000.50,
			StockQuantity: 5,
			Category:      model.CATEGORY_ENGINE,
			Dimensions:    model.Dimensions{Length: 2.5, Width: 1.0, Height: 1.2, Weight: 400},
			Manufacturer: model.Manufacturer{
				Name:    "NovaPropulsion",
				Country: "USA",
				Website: "https://novapropulsion.space",
			},
			Tags: []string{"engine", "ion", "propulsion"},
			Metadata: map[string]*model.Value{
				"Thrust": {DoubleValue: ptrFloat64(5000)},
			},
		},
		model.Part{
			UUID:          "d56dda39-3f77-4fed-bfe5-a614a6800bea",
			Name:          "Fusion Fuel Tank",
			Description:   "Tank designed for storing fusion fuel.",
			Price:         45000.0,
			StockQuantity: 10,
			Category:      model.CATEGORY_FUEL,
			Dimensions:    model.Dimensions{Length: 3.0, Width: 1.5, Height: 2.0, Weight: 800},
			Manufacturer: model.Manufacturer{
				Name:    "OrbitalFuelCo",
				Country: "Germany",
				Website: "https://orbitalfuelco.de",
			},
			Tags: []string{"fuel", "fusion"},
			Metadata: map[string]*model.Value{
				"CapacityLiters": {Int64Value: ptrInt64(15000)},
			},
		},
		model.Part{
			UUID:          "cb12b428-1ccf-4c1f-b5fd-53810446ea55",
			Name:          "Titanium Porthole",
			Description:   "Reinforced observation porthole made of titanium alloy.",
			Price:         7500.75,
			StockQuantity: 25,
			Category:      model.CATEGORY_PORTHOLE,
			Dimensions:    model.Dimensions{Length: 0.8, Width: 0.8, Height: 0.1, Weight: 30},
			Manufacturer: model.Manufacturer{
				Name:    "SkyView Windows",
				Country: "Japan",
				Website: "https://skyview.jp",
			},
			Tags: []string{"porthole", "window"},
			Metadata: map[string]*model.Value{
				"Material": {StringValue: ptrString("Titanium Alloy")},
			},
		},
		model.Part{
			UUID:          "5db4fa65-a65d-4234-8d91-00a0e024c4fc",
			Name:          "Stabilizing Wing MkII",
			Description:   "Advanced aerodynamic wing for stabilization.",
			Price:         32000,
			StockQuantity: 7,
			Category:      model.CATEGORY_WING,
			Dimensions:    model.Dimensions{Length: 4.0, Width: 0.3, Height: 1.2, Weight: 250},
			Manufacturer: model.Manufacturer{
				Name:    "Aerodyne Systems",
				Country: "Canada",
				Website: "https://aerodyne.ca",
			},
			Tags: []string{"wing", "stabilizer"},
			Metadata: map[string]*model.Value{
				"LiftCoefficient": {DoubleValue: ptrFloat64(1.8)},
			},
		},
		model.Part{
			UUID:          "199b521b-de12-4810-b91a-151d1e64fbc4",
			Name:          "Multi-fuel Converter",
			Description:   "Allows switching between multiple fuel types mid-flight.",
			Price:         15500.99,
			StockQuantity: 3,
			Category:      model.CATEGORY_FUEL,
			Dimensions:    model.Dimensions{Length: 1.2, Width: 0.9, Height: 0.6, Weight: 120},
			Manufacturer: model.Manufacturer{
				Name:    "FuelFlex Inc.",
				Country: "UK",
				Website: "https://fuelflex.co.uk",
			},
			Tags: []string{"fuel", "converter"},
			Metadata: map[string]*model.Value{
				"SupportedFuels": {StringValue: ptrString("Hydrogen, Methane, Ammonia")},
			},
		},
	}

	_, err := collection.InsertMany(ctx, parts)
	return err
}

func ptrString(s string) *string {
	return &s
}

func ptrInt64(i int64) *int64 {
	return &i
}

func ptrFloat64(f float64) *float64 {
	return &f
}
