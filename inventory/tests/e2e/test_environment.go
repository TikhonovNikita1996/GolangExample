package integration

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/model"
)

func (env *TestEnvironment) InsertTestPart(ctx context.Context) (string, error) {
	part := model.Part{
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
	}

	// Используем базу данных из переменной окружения MONGO_DATABASE
	databaseName := os.Getenv("MONGO_DATABASE")
	if databaseName == "" {
		databaseName = "inventory" // fallback значение
	}

	_, err := env.Mongo.Client().Database(databaseName).Collection(partsCollectionName).InsertOne(ctx, part)
	if err != nil {
		return "", err
	}

	return part.UUID, nil
}

func (env *TestEnvironment) GetTestSightingInfo() *model.Part {
	return &model.Part{
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
	}
}

func (env *TestEnvironment) ClearPartsCollection(ctx context.Context) error {
	// Используем базу данных из переменной окружения MONGO_DATABASE
	databaseName := os.Getenv("MONGO_DATABASE")
	if databaseName == "" {
		databaseName = "inventory" // fallback значение
	}

	_, err := env.Mongo.Client().Database(databaseName).Collection(partsCollectionName).DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	return nil
}

func ptrString(s string) *string {
	return &s
}
