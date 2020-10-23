package models

import (
	//"API/Chess"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type moves struct {
	number int
	white  string
	black  string
}
type player struct {
	Username string
	Rating   string
	Result   string
}

type MongoGames struct {
	URL   string
	Pgn   []moves
	White player
	Black player
}
type BlogPost struct {
	Title  string
	Topic  string
	Body   string
	Author string
	Images []string
}

// type BlogPost struct{
// 	ID primitive.ObjectID `bson:"_id,omitempty"`
// 	Title string `bson:"title,omitempty"`
// 	TLDR string	`bson:"title,omitempty"`
// 	Body string `bson:"title,omitempty"`
// 	Notes string `bson:"title,omitempty"`
// 	images []string `bson:"title,omitempty"`
// }

// func (db *MongoDb)RequestCreate(g chessapi.Game) string {
// 	return "Created" + g.URL
// }
// func (db *MongoDb)RequestUpdateByUrl(url string) string {
// 	return "Updated"
// }
// func (db *MongoDb)RequestReadByUrl(url string) string {
// 	return "Readed"
// }
// func(db *MongoDb) RequestDeleteByUrl(url string) bool {
// 	return true
// }
func (db *MongoDb) Create(b BlogPost) interface{} {
	client := db.DB
	quickstartDatabase := client.Database("blog")
	postsCollection := quickstartDatabase.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	blogResult, err := postsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: b.Title},
		{Key: "topic", Value: b.Topic},
		{Key: "author", Value: b.Author},
		{Key: "body", Value: b.Body},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blogResult.InsertedID)
	return blogResult.InsertedID
}
func (db *MongoDb) UpdateByID() int64 {
	return 40
}
type ReadStruct struct{
	posts []bson.M
}
func (db *MongoDb) ReadAll() /*BlogPost*/{
	client := db.DB
	quickstartDatabase := client.Database("blog")
	postsCollection := quickstartDatabase.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	cursor, err := postsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var posts []bson.M
	if err = cursor.All(ctx, &posts); err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts)
	

}
func (db *MongoDb) DeleteByID() bool {
	return true
}
