package chessapi

import (
	"net/http"
	"io/ioutil"
    "fmt"
    "encoding/json"
	
)

type player struct {
	Name string `json:"username"`
	Result string `json:"result"`
	Rating int `json:"rating"`
}
type Game struct {
	URL string `json:"URL"`
	White player `json:"white"`
	Black player `json:"black"`
}
type GameData struct {
	Games []Game `json:"games"`
}

func getGames(body []byte) (*GameData, error) {
    var g = new(GameData)
    err := json.Unmarshal(body, &g)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return g, err
}
func CallChess(url string) *GameData{
	res, err := http.Get(url)
    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }

    g, err := getGames([]byte(body))

    fmt.Println(g.Games[1].White)
    return g
}