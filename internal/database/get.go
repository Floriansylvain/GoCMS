package database

import (
	"context"
	"errors"
)

func GetDocuments(location Location, filter interface{}) ([][]byte, error) {
	client := GetNewClient()
	collection := client.Database(location.Database).Collection(location.Collection)
	defer client.Disconnect(context.TODO())

	var results [][]byte

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return results, errors.New("something is wrong with filter to find the document.")
	}
	for cursor.TryNext(context.TODO()) {
		results = append(results, cursor.Current)
	}
	return results, nil
}

func GetUniqueDocument(location Location, filter interface{}) ([]byte, error) {
	client := GetNewClient()
	collection := client.Database(location.Database).Collection(location.Collection)
	defer client.Disconnect(context.TODO())

	singleResult := collection.FindOne(context.TODO(), filter)
	if singleResult.Err() != nil {
		return nil, errors.New("no document was found, check the filter or the location.")
	}
	return singleResult.DecodeBytes()
}
