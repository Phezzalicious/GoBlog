package models

import(
	//"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"API/Chess"
	// "encoding/csv"
	"fmt"
	// "io"
	// "os"
	// "strings"
)

type game struct {
	URL string 
	WhiteName string
	WhiteResult string
	WhiteRating int
	BlackName string 
	BlackResult string
	BlackRating int
} 

type Games []game



func (db *MySqlDb)RequestCreate(g chessapi.Game) int64 {
	gg := game{g.URL,g.White.Name,g.White.Result,g.White.Rating,g.Black.Name,g.Black.Result,g.Black.Rating}
	Qs := fmt.Sprintf("Insert into Games(URL, WhiteName,WhiteResult,WhiteRating,BlackName,BlackResult,BlackRating) values(\"%s\",\"%s\",\"%s\",\"%d\",\"%s\",\"%s\",\"%d\")", gg.URL,gg.WhiteName, gg.WhiteResult, gg.WhiteRating, gg.BlackName, gg.BlackResult, gg.BlackRating)
	fmt.Println(Qs)
	res, err := db.DB.Exec(Qs)
	if err != nil{
		log.Fatal(err)
	}
	ra, _ := res.RowsAffected()
	id, _ := res.LastInsertId()

	log.Println("Rows Affected", ra, "Last inserted ID", id)
	return id
}

func (db *MySqlDb)RequestUpdateByUrl(url string) int64 { 
	return 5
}
func (db *MySqlDb)RequestReadByUrl(url string) { 

}
func (db *MySqlDb)RequestDeleteByUrl(url string) bool { 
	return true
}


	