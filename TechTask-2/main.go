package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	key := parts[2]

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}

	fmt.Println(key)
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/article/", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Articles = []Article{
		Article{ID: "1", Title: "Hello", Subtitle: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Subtitle: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}
