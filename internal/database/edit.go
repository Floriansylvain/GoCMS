package database

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type DocumentUpdate struct {
	Filter gin.H `json:"filter"`
	Update gin.H `json:"update"`
}

func EditDocument(location Location, jsons DocumentUpdate) (int64, error) {
	client := GetNewClient()
	collection := client.Database(location.Database).Collection(location.Collection)
	defer client.Disconnect(context.TODO())

	result, err := collection.UpdateOne(context.TODO(), jsons.Filter, gin.H{"$set": jsons.Update})
	if err != nil {
		return 0, err
	}
	if result.MatchedCount == 0 {
		return 0, errors.New("cannot find document matching filter.")
	}

	return result.ModifiedCount, nil
}
