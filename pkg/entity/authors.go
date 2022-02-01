package entity

import "github.com/test_kompas/news_app/internal/pkg"

type Authors struct {
	Name       string `json:"name" gorm:"type:varchar(100)"`
	Username   string `json:"username" gorm:"type:varchar(100);not null"`
	Password   string `json:"password" gorm:"type:varchar(100);not null"`
	BaseEntity `gorm:"embedded"`
}

type AuthorsService interface {
	FindAuthors(map[string]interface{}, *Pagination) ([]Authors, *pkg.Errors)
	AddAuthor(*Authors) (uint, *pkg.Errors)
	UpdateAuthor(uint, *Authors) *pkg.Errors
	DeleteAuthor(uint) *pkg.Errors
}

type AuthorsRepository interface {
	FindByName(string, *Pagination) ([]Authors, error)
	FindAll(*Pagination) ([]Authors, error)
	FindByUsername(string) (Authors, error)
	FindByID(uint) (Authors, error)
	AddAuthor(*Authors) error
	UpdateAuthor(uint, *Authors) error
	DeleteAuthor(uint) error
}

func (e *Authors) TableName() string {
	return "dtm_authors"
}
