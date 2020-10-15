package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"../controllers"
	"../modules"
	"../utils"
)

var mutex = &sync.Mutex{}

//HomePage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//Articles ...
func Articles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllArticles(w, r)
		return
	case "POST":
		postArticle(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	mutex.Lock()
	newarticle, err := modules.CreateNewArticle(reqBody)
	mutex.Unlock()
	if err {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cannot Save the Article"))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newarticle)
	}
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	limit := 5
	page, begin := utils.Pagination(r, limit)
	if begin >= len(modules.Articles) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Pages Ended"))
	} else {

		mutex.Lock()
		pages, TotalPages, results := utils.PagingArticle(page, begin, limit)
		mutex.Unlock()
		var next, prev int
		if page == pages {
			next = 1
			prev = page - 1
		} else {
			if page-1 < 0 {
				next = page + 1
				prev = 0
			} else {
				next = page + 1
				prev = page - 1
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Total    int               `json:"total"`
			Page     int               `json:"page"`
			Pages    int               `json:"pages"`
			NextPage int               `json:"nextpage"`
			PrevPage int               `json:"previouspage"`
			Articles []modules.Article `json:"docs"`
		}{
			Total:    TotalPages,
			Page:     page,
			Pages:    pages,
			NextPage: next,
			PrevPage: prev,
			Articles: results,
		})

	}

	fmt.Printf("Current Page: %d, Begin: %d\n", page, begin)
	fmt.Println("Endpoint Hit: returnAllArticles")
}

//GetArticleByID ...
func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	key := parts[2]

	mutex.Lock()
	article, isFound := controllers.GetArticleByID(key)
	defer mutex.Unlock()

	if isFound == true {
		json.NewEncoder(w).Encode(article)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Println(key)
	fmt.Println("Endpoint Hit: homePage")
	return
}

//SearchArticlesByKey ...
func SearchArticlesByKey(w http.ResponseWriter, r *http.Request) {

	key := r.FormValue("q")
	mutex.Lock()

	FoundArticles := controllers.SearchByTitleSubtitleAndContent(key)
	defer mutex.Unlock()

	json.NewEncoder(w).Encode(FoundArticles)
	fmt.Println("Endpoint Hit: Search Query")
	return
}
