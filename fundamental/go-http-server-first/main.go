package main

import (
	"html/template"
	"net/http"
	"time"
	"os"
)

type Source struct {
	ID   interface{} `json:"id"`
	Name string      `json:"name"`
}
type Article struct {
	Source      Source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}
type Results struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}
type Search struct {
	SearchKey  string
	NextPage   int
	TotalPages int
	Results    Results
}

var apiKey *string
var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	// apiKey = flag.String("apiKey", "", "Newsapi.org access key")
	// flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}