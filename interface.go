package mongow

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient interface {
	Disconnect(ctx context.Context)
	DB(name string) (MongoDB, error)
}

type MongoDB interface {
	Collection(name string) (*mongo.Collection, error)
}
