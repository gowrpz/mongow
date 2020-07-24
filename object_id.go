package mongow

import (
	"github.com/goutlz/errz"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func NewObjectIdFromString(id string) (interface{}, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errz.Wrapf(err, "Failed to create object ID from string %s", id)
	}

	return objectId, nil
}

func ConvertInsertedIds(res *mongo.InsertManyResult) ([]string, error) {
	var resIds []string

	for _, id := range res.InsertedIDs {
		objectId, ok := id.(primitive.ObjectID)
		if ok {
			resIds = append(resIds, objectId.Hex())
			continue
		}

		el, ok := id.(*bsonx.Elem)
		if !ok {
			return []string{}, errz.Newf("ObjectId type is unexpected: %T", id)
		}

		idStr, ok := el.Value.StringValueOK()
		if ok {
			resIds = append(resIds, idStr)
			continue
		}

		idObj, ok := el.Value.ObjectIDOK()
		if ok {
			resIds = append(resIds, idObj.Hex())
			continue
		}

		return []string{}, errz.Newf("Inserted ObjectId type is unexpected: %T. Value-type: %+v", id, el.Value.Type())
	}

	return resIds, nil
}
