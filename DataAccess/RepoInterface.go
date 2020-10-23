package models
import("API/Chess"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/bson/primitive"
    "log"
    //"fmt"
    "go.mongodb.org/mongo-driver/mongo/options")

type GameStore interface {
	 
    RequestCreate(chessapi.Game) int64
    RequestUpdateByUrl(url string) int64
    RequestReadByUrl(url string) 
	RequestDeleteByUrl(url string) bool
}
type BlogStore interface { 
    Create(b BlogPost) interface{}
    UpdateByID() int64
    ReadAll() //BlogPost
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
	
		return nil, err
		
	}
	
    return &MySqlDb{db2}, nil
}

type MongoDb struct{
	DB *mongo.Client
}
func NewMongoDB(dataSourceName string) (*MongoDb, error){

    client, err := mongo.NewClient(options.Client().ApplyURI(dataSourceName))
	
	if err != nil{
		log.Fatal(err)
	}
	
    return &MongoDb{client}, nil
	
}



