package mongow

import (
	"context"
	"time"

	"github.com/goutlz/errz"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClientWrap struct {
	client *mongo.Client
}

func (mc *mongoClientWrap) Disconnect(ctx context.Context) {
	mc.client.Disconnect(ctx)
}

func (mc *mongoClientWrap) DB(name string) (MongoDB, error) {
	db := mc.client.Database(name)
	if db == nil {
		return nil, errz.New("DB is nil")
	}

	return &databaseWrap{
		db: db,
	}, nil
}

func NewClient(ctx context.Context, connStr string, timeout time.Duration) (MongoClient, error) {
	clientOpts := options.Client().SetConnectTimeout(timeout)
	clientOpts = clientOpts.ApplyURI(connStr)

	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		return nil, errz.Wrap(err, "Failed to make a new mongodb client")
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, errz.Wrap(err, "Failed to connect to mongodb client")
	}

	return &mongoClientWrap{
		client: client,
	}, nil
}
