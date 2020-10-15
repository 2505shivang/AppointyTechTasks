package controllers

import (
	"strings"

	"../modules"
)

//GetArticleByID ...
func GetArticleByID(ID string) (modules.Article, bool) {

	for _, article := range modules.Articles {
		if article.ID == ID {
			return article, true
			//json.NewEncoder(w).Encode(article)
		}
	}

	var a modules.Article
	return a, false
}

//GetArticleBySubtitle ...
func GetArticleBySubtitle(key string) []modules.Article {

	var FoundArticles []modules.Article
	for _, article := range modules.Articles {
		inSubtitle := strings.Contains(strings.ToLower(article.Subtitle), strings.ToLower(key))
		if inSubtitle == true {
			FoundArticles = append(FoundArticles, article)
		}
	}
	return FoundArticles
}

//GetArticleByTitle ...
func GetArticleByTitle(key string) []modules.Article {

	var FoundArticles []modules.Article
	for _, article := range modules.Articles {
		intitle := strings.Contains(strings.ToLower(article.Title), strings.ToLower(key))
		if intitle == true {
			FoundArticles = append(FoundArticles, article)
		}
	}

	return FoundArticles

}

//GetArticleByContent ...
func GetArticleByContent(key string) []modules.Article {

	var FoundArticles []modules.Article
	for _, article := range modules.Articles {
		inContent := strings.Contains(strings.ToLower(article.Content), strings.ToLower(key))
		if inContent == true {
			FoundArticles = append(FoundArticles, article)
		}
	}
	return FoundArticles
}

//SearchByTitleSubtitleAndContent ...
func SearchByTitleSubtitleAndContent(key string) []modules.Article {
	FoundArticlesByTitle := GetArticleByTitle(key)
	FoundArticlesBySubTitle := GetArticleBySubtitle(key)
	FoundArticlesByContent := GetArticleByContent(key)

	FoundArticles := []modules.Article{}
	FoundArticles = append(FoundArticlesByTitle, FoundArticlesBySubTitle...)
	FoundArticles = append(FoundArticles, FoundArticlesByContent...)

	return FoundArticles
}
