package mongow

import (
	"github.com/goutlz/errz"

	"go.mongodb.org/mongo-driver/mongo"
)

type databaseWrap struct {
	db *mongo.Database
}

func (d *databaseWrap) Collection(name string) (*mongo.Collection, error) {
	coll := d.db.Collection(name)
	if coll == nil {
		return nil, errz.New("Collection is nil")
	}

	return coll, nil
}
