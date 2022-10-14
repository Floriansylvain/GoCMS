package internal

import (
	"context"

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
	defer client.Disconnect(context.TODO())

	coll := client.Database(location.Database).Collection(location.Collection)
	_, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}

	return nil
}

// func getDocument(database string, collection string) {
// }
