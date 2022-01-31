package repository

import (
	"fmt"

	"github.com/test_kompas/news_app/pkg/entity"
	"gorm.io/gorm"
)

type AuthorsRepository struct {
	baseRepository
}

func NewAuthorsRepository(db *gorm.DB) *AuthorsRepository {
	tableName := (&entity.Authors{}).TableName()
	db.AutoMigrate(&entity.Authors{})
	baseRepository := New(db, tableName)
	return &AuthorsRepository{baseRepository}
}

func (repo *AuthorsRepository) FindByName(name string, options *entity.Pagination) (results []entity.Authors, err error) {
	results = make([]entity.Authors, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	q := fmt.Sprintf("%%%s%%", name)
	err = tx.Where("name LIKE ?", q).Order("id").Find(&results).Error
	return
}

func (repo *AuthorsRepository) FindAll(options *entity.Pagination) (results []entity.Authors, err error) {
	results = make([]entity.Authors, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	err = tx.Table(repo.tableName).Order("id").Find(&results).Error
	return
}

func (repo *AuthorsRepository) FindByUsername(username string) (data entity.Authors, err error) {
	err = repo.db.Where("username = ?", username).Order("id").Find(&data).Error
	return
}

func (repo *AuthorsRepository) FindByID(id int) (data entity.Authors, err error) {
	err = repo.db.Where("id = ?", id).Order("id").Find(&data).Error
	return
}

func (repo *AuthorsRepository) AddAuthor(author *entity.Authors) (err error) {
	err = repo.DBInsert(author)
	return
}

func (repo *AuthorsRepository) UpdateAuthor(authorId int, author *entity.Authors) (err error) {
	err = repo.UpdateByID(authorId, author)
	return
}

func (repo *AuthorsRepository) DeleteAuthor(authorId int) (err error) {
	author := entity.Authors{}
	author.SetID(authorId)
	err = repo.DeleteByID(&author)
	return
}
