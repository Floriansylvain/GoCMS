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

func getCollection(location Location) *mongo.Collection {
	client := getNewClient()
	defer client.Disconnect(context.TODO())
	return client.Database(location.Database).Collection(location.Collection)
}

func pushDocument(location Location, document interface{}) error {
	collection := getCollection(location)
	_, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}
	return nil
}

func getDocuments(location Location, filter interface{}) ([][]byte, error) {
	collection := getCollection(location)
	var results [][]byte

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return results, errors.New("something is wrong with filter to find document")
	}
	for cursor.TryNext(context.TODO()) {
		results = append(results, cursor.Current)
	}
	return results, nil
}
