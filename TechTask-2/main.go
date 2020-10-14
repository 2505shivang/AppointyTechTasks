package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func articles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		returnAllArticles(w, r)
		return
	case "POST":
		createNewArticle(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", articles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Articles = []Article{
		Article{ID: "1", Title: "Hello", Subtitle: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Subtitle: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}
