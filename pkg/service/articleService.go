package service

import (
	"fmt"
	"regexp"
	"time"

	"github.com/test_kompas/news_app/internal/pkg"
	"github.com/test_kompas/news_app/pkg/entity"
)

type ArticleService struct {
	repo entity.ArticleRepository
}

func NewArticleService(repo entity.ArticleRepository) *ArticleService {
	return &ArticleService{repo}
}

func (srv *ArticleService) FindArticles(query map[string]interface{}, pagination *entity.Pagination) (results []entity.Article, err *pkg.Errors) {
	results = make([]entity.Article, 0)

	if v, ok := query["id"]; ok {
		if value, ok := v.(uint); ok && value != 0 {

			if data, e := srv.repo.FindByID(value); e != nil {
				err = pkg.NewError(
					fmt.Sprintf("Could not retrieve data : %s", e.Error()),
					500,
				)
			} else {
				results = append(results, data)
			}
			return
		}
	}

	if v, ok := query["title"]; ok {
		if data, ok := v.(string); ok && data != "" {
			var e error
			if results, e = srv.repo.FindByTitle(data, pagination); e != nil {
				err = pkg.NewError(
					fmt.Sprintf("Could not retrieve data : %s", e.Error()),
					500,
				)
			}
			return
		}
	}

	var e error
	if results, e = srv.repo.FindAll(pagination); e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Could not retrieve data : %s", e.Error()),
			500,
		)
	}
	return
}

func (srv *ArticleService) AddArticle(article *entity.Article, authorId uint) (err *pkg.Errors) {

	if article.Title == "" || article.Body == "" {
		err = pkg.NewError("article title and content must not be empty", 400)
		return
	}

	getArticles, e := srv.repo.FindByTitle(article.Title, &entity.Pagination{})
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Could not retreive article data : %s", e.Error()),
			500,
		)
		return
	}
	if len(getArticles) > 0 {
		err = pkg.NewError(
			fmt.Sprintf("article \"%s\" already exists", article.Title),
			404,
		)
		return
	}

	if article.Status != "PUBLISHED" {
		article.Status = "UNPUBLISHED"
	}
	article.ReleasedDate = time.Now().Format("2006-01-02 15:04:05")
	article.AuthorID = authorId

	re, e := regexp.Compile(`<(.*?)>`)
	article.Content = re.ReplaceAllString(article.Body, " ")

	if e := srv.repo.AddArticle(article); e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Add new article failed : %s", e.Error()),
			500,
		)
	}
	return
}
