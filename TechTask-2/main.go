package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//Data Structure

type Article struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	Subtitle          string    `json:"subtitle"`
	Content           string    `json:"content"`
	CreationTimestamp time.Time `json:"timestamp"`
}

var Articles []Article

//handlers

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

	article.CreationTimestamp = time.Now()

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

	w.WriteHeader(http.StatusNotFound)
	fmt.Println(key)
	fmt.Println("Endpoint Hit: homePage")
	return
}

//search
func searchquery(w http.ResponseWriter, r *http.Request) {

	v := r.FormValue("q")

	var FoundArticles []Article

	for _, article := range Articles {

		intitle := strings.Contains(strings.ToLower(article.Title), strings.ToLower(v))
		inSubtitle := strings.Contains(strings.ToLower(article.Subtitle), strings.ToLower(v))
		inContent := strings.Contains(strings.ToLower(article.Content), strings.ToLower(v))

		fmt.Println(inSubtitle)
		fmt.Println(intitle)
		fmt.Println(inContent)

		if intitle == true || inSubtitle == true || inContent == true {
			FoundArticles = append(FoundArticles, article)
		}
	}

	json.NewEncoder(w).Encode(FoundArticles)
	return
}

//server

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/articles/", returnSingleArticle)
	http.HandleFunc("/articles/search", searchquery)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//main function

func main() {
	Articles = []Article{
		Article{ID: "1", Title: "Hello", Subtitle: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Subtitle: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}
