package mongow

import (
	"github.com/goutlz/errz"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func NewObjectIdFromString(id string) (interface{}, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errz.Wrapf(err, "Failed to create object ID from string %s", id)
	}

	return objectId, nil
}

func ConvertObjectIdToStringValue(id interface{}) (string, error) {
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
