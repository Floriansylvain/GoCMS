package database

import (
	"context"
	"errors"
)

func DeleteDocument(location Location, filter interface{}) (int64, error) {
	client := GetNewClient()
	collection := client.Database(location.Database).Collection(location.Collection)
	defer client.Disconnect(context.TODO())

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, errors.New("something is wrong with filter to delete the document.")
	}

	return result.DeletedCount, nil
}
