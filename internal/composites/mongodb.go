package composites

import (
	"CRMka/internal/config"
	"CRMka/pkg/client/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBComposite struct {
	db *mongo.Database
}

func NewMongoDBComposite(ctx context.Context, cfg *config.Config) *MongoDBComposite {
	c := cfg.MongoDB
	client, err := mongodb.NewClient(ctx, c.Host, c.Port, c.Username, c.Password, c.Database, c.AuthDB)
	if err != nil {
		panic(err)
	}
	return &MongoDBComposite{db: client}
}
