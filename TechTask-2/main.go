package main

import (
	"log"
	"net/http"

	"./go-Packages/api"
	"./go-Packages/modules"
)

//server

func handleRequests() {
	http.HandleFunc("/", api.HomePage)
	http.HandleFunc("/articles", api.Articles)
	http.HandleFunc("/articles/", api.GetArticleByID)
	http.HandleFunc("/articles/search", api.SearchArticlesByKey)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//main function

func main() {

	modules.Articles = []modules.Article{
		modules.Article{ID: 1, Title: "Hello", Subtitle: "Article Description", Content: "Article Content"},
		modules.Article{ID: 2, Title: "Hello 2", Subtitle: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}
