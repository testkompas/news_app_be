package entity

import (
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/test_kompas/news_app/internal/pkg"
)

type Authentication struct {
	AuthorID   int     `json:"author_id" gorm:"type:int(10);not null"`
	Token      string  `json:"token" gorm:"type:text;not null"`
	IssuedAt   string  `json:"issued_at" gorm:"type:varchar(20);not null"`
	ExpiryTime int32   `json:"expiry_time" gorm:"type:int(10);not null"`
	Authors    Authors `json:"-" gorm:"foreignKey:AuthorID"`
	BaseEntity `gorm:"embedded"`
}

type AuthenticationService interface {
	AuthorLogin(*Authors) (string, *pkg.Errors)
	Authorize(string) (jwt.Token, *pkg.Errors)
}

type AuthenticationRepository interface {
	FindBytoken(string) (Authentication, error)
	FindAll(*Pagination) ([]Authentication, error)
	FindByID(int) (Authentication, error)
	AddToken(*Authentication) error
	DeleteToken(int) error
}

func (e *Authentication) TableName() string {
	return "dth_authentication"
}
