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
		strId, err := convertInterfaceObjectIdToStringValue(id)
		if err != nil {
			return nil, errz.Wrap(err, "Failed to convert interface object ID to string value")
		}

		resIds = append(resIds, strId)
	}

	return resIds, nil
}

func ConverInsertedId(res *mongo.InsertOneResult) (string, error) {
	id := res.InsertedID
	strId, err := convertInterfaceObjectIdToStringValue(id)
	if err != nil {
		return "", errz.Wrap(err, "Failed to convert interface object ID to string value")
	}

	return strId, nil
}

func convertInterfaceObjectIdToStringValue(id interface{}) (string, error) {
	objectId, ok := id.(primitive.ObjectID)
	if ok {
		return objectId.Hex(), nil
	}

	el, ok := id.(*bsonx.Elem)
	if !ok {
		return "", errz.Newf("ObjectId type is unexpected: %T", id)
	}

	idStr, ok := el.Value.StringValueOK()
	if ok {
		return idStr, nil
	}

	idObj, ok := el.Value.ObjectIDOK()
	if ok {
		return idObj.Hex(), nil
	}

	return "", errz.Newf("ObjectId type is unexpected: %T. Value-type: %+v", id, el.Value.Type())
}
