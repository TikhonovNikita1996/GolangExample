package part

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/inventory/internal/helpers"
)

type repository struct {
	mu sync.RWMutex

	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *repository {
	collection := db.Collection("parts")

	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "uuid", Value: 1}},
			Options: options.Index().SetUnique(false),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		panic(err)
	}

	count, _ := collection.CountDocuments(ctx, bson.D{})
	if count == 0 {
		helpers.SeedParts(ctx, collection)
	}
	return &repository{
		collection: collection,
	}
}
