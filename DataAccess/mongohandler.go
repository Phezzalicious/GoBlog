package models
import (
	//"API/Chess"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type moves struct{
	number int
	white string
	black string
}
type player struct{
	Username string
	Rating string
	Result string
}

type MongoGames struct{
	URL string
	Pgn []moves
	White player
	Black player
}
type BlogPost struct{
	Title string 
	TLDR string
	Body string 
	Notes string
	images []string
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
func (db *MongoDb)Create() int64 {
	return 40
}
func (db *MongoDb)UpdateByID() int64 { 
	return 40
}
func (db *MongoDb)ReadByID(id primitive.ObjectID) { 
	
}
func(db *MongoDb) DeleteByID() bool { 
	return true
}