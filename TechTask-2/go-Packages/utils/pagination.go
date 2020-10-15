package utils

import (
	"net/http"
	"strconv"

	"../modules"
)

//Pagination ...
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

//PagingArticle ...
func PagingArticle(page, begin, limit int) (int, int, []modules.Article) {
	total := len(modules.Articles)
	pages := (total / limit)
	var results []modules.Article

	if (total % limit) != 0 {
		pages++
	}

	if page*limit > total {
		results = modules.Articles[begin:total]
	} else {
		results = modules.Articles[begin : page*limit]
	}

	return pages, total, results

}
