package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Articles = []Article{
		Article{ID: "1", Title: "Hello", Subtitle: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Subtitle: "Article Description", Content: "Article Content"},
	}
	//handleRequests()
}
