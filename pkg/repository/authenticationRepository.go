package repository

import (
	"github.com/test_kompas/news_app/pkg/entity"
	"gorm.io/gorm"
)

type AuthenticationRepository struct {
	baseRepository
}

func NewAuthenticationRepository(db *gorm.DB) *AuthenticationRepository {
	tableName := (&entity.Authentication{}).TableName()
	db.AutoMigrate(&entity.Authentication{})
	baseRepository := New(db, tableName)
	return &AuthenticationRepository{baseRepository}
}

func (repo *AuthenticationRepository) FindBytoken(token string) (result entity.Authentication, err error) {
	err = repo.db.Where("token = ?", token).Order("id").Find(&result).Error
	return
}

func (repo *AuthenticationRepository) FindAll(options *entity.Pagination) (results []entity.Authentication, err error) {
	results = make([]entity.Authentication, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	err = tx.Table(repo.tableName).Order("id").Find(&results).Error
	return
}

func (repo *AuthenticationRepository) FindByID(id uint) (data entity.Authentication, err error) {
	err = repo.db.Where("id = ?", id).Order("id").Find(&data).Error
	return
}

func (repo *AuthenticationRepository) AddToken(authentication *entity.Authentication) (err error) {
	err = repo.DBInsert(authentication)
	return
}

func (repo *AuthenticationRepository) DeleteToken(tokenId uint) (err error) {
	authentication := entity.Authentication{}
	authentication.SetID(tokenId)
	err = repo.DeleteByID(&authentication)
	return
}
