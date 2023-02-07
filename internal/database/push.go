package database

import "context"

func PushDocument(location Location, document interface{}) error {
	client := GetNewClient()
	collection := client.Database(location.Database).Collection(location.Collection)
	defer client.Disconnect(context.TODO())

	_, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}
	return nil
}
