package internal

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Location struct {
	Database   string
	Collection string
}

func getNewClient() *mongo.Client {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://db:27017/"))
	if err != nil {
		panic(err)
	}
	return client
}

func pushDocument(location Location, document interface{}) error {
	client := getNewClient()
	collection := client.Database(location.Database).Collection(location.Collection)
	defer client.Disconnect(context.TODO())

	_, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}
	return nil
}

func getDocuments(location Location, filter interface{}) ([][]byte, error) {
	client := getNewClient()
	collection := client.Database(location.Database).Collection(location.Collection)
	defer client.Disconnect(context.TODO())

	var results [][]byte

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return results, errors.New("something is wrong with filter to find the document")
	}
	for cursor.TryNext(context.TODO()) {
		results = append(results, cursor.Current)
	}
	return results, nil
}

func getUniqueDocument(location Location, filter interface{}) ([]byte, error) {
	client := getNewClient()
	collection := client.Database(location.Database).Collection(location.Collection)
	defer client.Disconnect(context.TODO())

	singleResult := collection.FindOne(context.TODO(), filter)
	if singleResult.Err() != nil {
		return nil, errors.New("no document was found, check the filter or the location.")
	}
	return singleResult.DecodeBytes()
}

func deleteDocument(location Location, filter interface{}) (int64, error) {
	client := getNewClient()
	collection := client.Database(location.Database).Collection(location.Collection)
	defer client.Disconnect(context.TODO())

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, errors.New("something is wrong with filter to delete the document")
	}

	return result.DeletedCount, nil
}
