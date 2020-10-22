package main

import (
	"fmt"
	//"html/template"
	"log"
	"net/http"
	"API/DataAccess"
	"html/template"
	"API/Chess"

)

type Env struct {
	db models.GameStore
	blogDB models.BlogStore
}


func main() {
	fmt.Println("YOYOYO")
	db, err := models.NewMySqlDB("phelps:1995@/GoLang")
    if err != nil {
       log.Panic(err)
	}
	fmt.Println("YOYOYO")
	
	db2, _ := models.NewMongoDB("mongodb+srv://phelps:1995@cluster0.ddb3v.mongodb.net/blog?retryWrites=true&w=majority")
	env := &Env{db,db2}
	
	url := "https://api.chess.com/pub/player/Phezzalicious/games/2020/03"
	seed := chessapi.CallChess(url)

	//id := env.db.RequestCreate(seed.Games[2])
	fmt.Println(len(seed.Games))


	//fmt.Println("MY ADDED ID: " + fmt.Sprint(id))
	
	http.HandleFunc("/", env.IndexHandler)
	http.HandleFunc("/glu", env.GluHandler)
	http.HandleFunc("/agg", env.AggHandler)
	http.HandleFunc("/blog",env.BlogHomeHandler)
	http.HandleFunc("/createBlog", env.CreateBlogHandler)
	http.ListenAndServe(":8000", nil)
}


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}
type cBlogStruct struct{
	Title string
	Heading string
	Message string
}
func(env *Env) CreateBlogHandler(w http.ResponseWriter, r *http.Request){
	c := cBlogStruct{"Create a blog today!", "Blog form", "Make a blog"}
	err := tpl.ExecuteTemplate(w,"createblog.gohtml",c)
	if err != nil {
		log.Fatalln(err)
	}
}

type blogStruct struct{
	Title string
	Heading string
	Message string
	BlogTitle string
	Author string
	Body string
	Topic string
}
func (env *Env) BlogHomeHandler(w http.ResponseWriter, r *http.Request){
	p := blogStruct{"Blog","Blog Posts","Posts for the blog","Title","Author","Body","Topic"}
	err := tpl.ExecuteTemplate(w,"blog.gohtml",p)
	if err != nil {
		log.Fatalln(err)
	}
}


type aggStruct struct{
	Title string
	Username string
	Heading string
	Message string
	Games []chessapi.Game	
}

func (env *Env) AggHandler(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("uname")
	date := r.FormValue("date")
	url := "https://api.chess.com/pub/player/"+username+"/games/"+date
	g := chessapi.CallChess(url)
	p:= aggStruct{"Aggregated Game Data",username,fmt.Sprintf("Games for %s", username), fmt.Sprintf("During %s", date),g.Games}
	err := tpl.ExecuteTemplate(w,"agg.gohtml",p)
	if err != nil {
		log.Fatalln(err)
	}	
}

type gluStruct struct {
	Title string
	Heading string
	Message string
}
func (env *Env) GluHandler(w http.ResponseWriter, r *http.Request) {
	p:= gluStruct{"GameLookUp","Game Search", "Pick a username and date!"}
	err := tpl.ExecuteTemplate(w,"glu.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}	
}

type indexStruct struct {
	Title string
	GameOD chessapi.Game
	Heading string
	Message string
}

func (env *Env) IndexHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.chess.com/pub/player/phezzalicious/games/2020/02"
	g := chessapi.CallChess(url)
	p := indexStruct{Title: "Home", GameOD: g.Games[0], Heading: "Chess.com API Interations",Message: "More to come!"}

	//C:\Code\GO\MyBlog\Site\index.html
	err := tpl.ExecuteTemplate(w, "index.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
	
	
}






