package entity

import (
	"github.com/test_kompas/news_app/internal/pkg"
)

type Article struct {
	Title        string  `json:"title" gorm:"type:varchar(50);not null"`
	Body         string  `json:"body" gorm:"type:text;not null"`
	ReleasedDate string  `json:"released_date" gorm:"type:varchar(20);not null"`
	Status       string  `json:"status" gorm:"type:varchar(11);not null"`
	AuthorID     int     `json:"author_id" gorm:"type:int(10);not null"`
	Author       Authors `json:"-" gorm:"foreignKey:AuthorID"`
	BaseEntity   `gorm:"embedded"`
}

type ArticleService interface {
	AddArticle(*Article, int) *pkg.Errors
	FindArticles(map[string]interface{}, *Pagination) ([]Article, *pkg.Errors)
}

type ArticleRepository interface {
	FindByTitle(string, *Pagination) ([]Article, error)
	FindAll(*Pagination) ([]Article, error)
	FindByID(int) (Article, error)
	AddArticle(*Article) error
	DeleteArticle(int) error
}

func (e *Article) TableName() string {
	return "dtm_articles"
}
