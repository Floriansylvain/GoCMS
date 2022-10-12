package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gotest/internal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic("Error loading .env file.")
	}

	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	r.GET("/ping", internal.Ping)
	r.POST("/add-article", internal.AddArticle)

	r.Run()
}

// type album struct {
// 	ID    primitive.ObjectID `json:"id"`
// 	Pouet int32              `json:"number"`
// }

// func pouet(c *gin.Context) {
// 	const uri = "mongodb://db:27017/"
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer client.Disconnect(context.TODO())

// 	coll := client.Database("gohcms").Collection("posts")
// 	found, _ := coll.Find(context.TODO(), bson.D{})

// 	for found.Next(context.TODO()) {
// 		// fmt.Println(found.Current.String())
// 		var pouet album
// 		err := found.Decode(&pouet)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(pouet)
// 		fmt.Println(bson.Marshal(pouet))
// 	}

// 	c.JSON(http.StatusOK, gin.H{})
// }
