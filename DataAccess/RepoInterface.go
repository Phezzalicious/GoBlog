package models
import("API/Chess"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "log"
    "context"
    "time"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo/options")

type GameStore interface {
	 
    RequestCreate(chessapi.Game) int64
    RequestUpdateByUrl(url string) int64
    RequestReadByUrl(url string) 
	RequestDeleteByUrl(url string) bool
}
type BlogStore interface { 
    Create() int64
    UpdateByID() int64
    ReadByID(id primitive.ObjectID) 
	DeleteByID() bool
}

type MySqlDb struct{
	DB *sql.DB
}
func NewMySqlDB(dataSourceName string) (*MySqlDb, error) {
	
	db, err := sql.Open("mysql", dataSourceName)
	
	db.Close()
	db2, err2 := sql.Open("mysql", dataSourceName)
    if err2 != nil {
		
		return nil, err
		
    }
    if err2   = db.Ping(); err != nil {
		fmt.Println("HEYHEYHEY3")
		return nil, err
		
	}
	
    return &MySqlDb{db2}, nil
}

type MongoDb struct{
	db *mongo.Client
}
func NewMongoDB(dataSourceName string) (*MongoDb, error){

    client, err := mongo.NewClient(options.Client().ApplyURI(dataSourceName))
	
	if err != nil{
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}
    defer client.Disconnect(ctx)
    return &MongoDb{client}, nil
	//quickstartDatabase := client.Database("blog")
	//postsCollection := quickstartDatabase.Collection("posts")

	// blogResult, err := postsCollection.InsertOne(ctx, bson.D{
	// 	{Key: "title", Value: "My journey in programming"},
	// 	{Key: "topic", Value: "Code"},
	// 	{Key: "Author", Value: "Phelps Merrell"},
	// })
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(blogResult.InsertedID)	
}



