package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
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

var mutex = &sync.Mutex{}

//Pagination

func Pagination(r *http.Request, limit int) (int, int) {
	keys := r.URL.Query()
	if keys.Get("page") == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(keys.Get("page"))
	if page < 1 {
		return 1, 0
	}
	begin := (limit * page) - limit
	return page, begin
}

//handlers

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	limit := 5
	page, begin := Pagination(r, limit)

	var results []Article
	total := len(Articles)
	pages := (total / limit)
	if (total % limit) != 0 {
		pages++
	}
	mutex.Lock()
	if page*limit > total {
		results = Articles[begin:total]
	} else {
		results = Articles[begin : page*limit]
	}
	mutex.Unlock()

	fmt.Printf("Current Page: %d, Begin: %d\n", page, begin)
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(struct {
		Total    int       `json:"total"`
		Page     int       `json:"page"`
		Pages    int       `json:"pages"`
		NextPage int       `json:"nextpage"`
		PrevPage int       `json:"previouspage"`
		Articles []Article `json:"docs"`
	}{
		Total:    total,
		Page:     page,
		Pages:    pages,
		NextPage: page + 1,
		PrevPage: page - 1,
		Articles: results,
	})
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	mutex.Lock()
	article.CreationTimestamp = time.Now()
	Articles = append(Articles, article)
	defer mutex.Unlock()
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

	mutex.Lock()
	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
	defer mutex.Unlock()

	w.WriteHeader(http.StatusNotFound)
	fmt.Println(key)
	fmt.Println("Endpoint Hit: homePage")
	return
}

//search
func searchquery(w http.ResponseWriter, r *http.Request) {

	v := r.FormValue("q")

	var FoundArticles []Article

	mutex.Lock()
	for _, article := range Articles {

		intitle := strings.Contains(strings.ToLower(article.Title), strings.ToLower(v))
		inSubtitle := strings.Contains(strings.ToLower(article.Subtitle), strings.ToLower(v))
		inContent := strings.Contains(strings.ToLower(article.Content), strings.ToLower(v))

		if intitle == true || inSubtitle == true || inContent == true {
			FoundArticles = append(FoundArticles, article)
		}
	}
	defer mutex.Unlock()

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
