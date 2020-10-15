package modules

import (
	"encoding/json"
	"time"
)

//Article Schema ...
type Article struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	Subtitle          string    `json:"subtitle"`
	Content           string    `json:"content"`
	CreationTimestamp time.Time `json:"timestamp"`
}

//Articles ...
var Articles []Article

//CreateNewArticle ...
func CreateNewArticle(reqBody []byte) (Article, bool) {
	var article Article
	err := json.Unmarshal(reqBody, &article)
	if err != nil {
		return article, true
	}
	article.CreationTimestamp = time.Now()
	Articles = append(Articles, article)
	return article, false
}
