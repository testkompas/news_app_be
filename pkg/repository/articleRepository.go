package repository

import (
	"fmt"

	"github.com/test_kompas/news_app/pkg/entity"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	baseRepository
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	tableName := (&entity.Article{}).TableName()
	db.AutoMigrate(&entity.Article{})
	baseRepository := New(db, tableName)
	return &ArticleRepository{baseRepository}
}

func (repo *ArticleRepository) FindByTitle(title string, options *entity.Pagination) (result []entity.Article, err error) {
	result = make([]entity.Article, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	title = fmt.Sprintf("%%%s%%", title)
	err = tx.Where("title LIKE ?", title).Order("id").Find(&result).Error
	return
}

func (repo *ArticleRepository) FindAll(options *entity.Pagination) (results []entity.Article, err error) {
	results = make([]entity.Article, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	err = tx.Table(repo.tableName).Order("id").Find(&results).Error
	return
}

func (repo *ArticleRepository) FindByID(id int) (data entity.Article, err error) {
	err = repo.db.Where("id = ?", id).Order("id").Find(&data).Error
	return
}

func (repo *ArticleRepository) AddArticle(article *entity.Article) (err error) {
	err = repo.DBInsert(article)
	return
}

func (repo *ArticleRepository) DeleteArticle(articleId int) (err error) {
	Article := entity.Article{}
	Article.SetID(articleId)
	err = repo.DeleteByID(&Article)
	return
}
