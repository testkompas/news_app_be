package service

import (
	"fmt"

	"github.com/test_kompas/news_app/internal/pkg"
	"github.com/test_kompas/news_app/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type AuthorsService struct {
	repo entity.AuthorsRepository
}

func NewAuthorsService(repo entity.AuthorsRepository) *AuthorsService {
	return &AuthorsService{repo}
}

func (srv *AuthorsService) FindAuthors(query map[string]interface{}, pagination *entity.Pagination) (results []entity.Authors, err *pkg.Errors) {
	results = make([]entity.Authors, 0)

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

	if v, ok := query["name"]; ok {
		if data, ok := v.(string); ok && data != "" {
			var e error
			if results, e = srv.repo.FindByName(data, pagination); e != nil {
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

func (srv *AuthorsService) AddAuthor(author *entity.Authors) (authorId uint, err *pkg.Errors) {

	if author.Username == "" || author.Password == "" {
		err = pkg.NewError("authorname and password field is required", 400)
		return
	}
	getAuthor, e := srv.repo.FindByUsername(author.Username)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Could not retreive author data : %s", e.Error()),
			500,
		)
		return
	}
	if getAuthor.ID != 0 {
		err = pkg.NewError(
			fmt.Sprintf("Author \"%s\" already exists", author.Username),
			404,
		)
		return
	}

	passwordHash, e := srv.hashPassword(author.Password)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Password hash failed : %s", e.Error()),
			500,
		)
		return
	}
	author.Password = string(passwordHash[:])

	if e := srv.repo.AddAuthor(author); e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Add new author failed : %s", e.Error()),
			500,
		)
		return
	}
	authorId = author.ID
	return
}

func (srv *AuthorsService) UpdateAuthor(authorId uint, author *entity.Authors) (err *pkg.Errors) {

	if e := srv.repo.UpdateAuthor(authorId, author); e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Update author failed : %s", e.Error()),
			500,
		)
	}
	return
}

func (srv *AuthorsService) DeleteAuthor(authorId uint) (err *pkg.Errors) {

	if e := srv.repo.DeleteAuthor(authorId); e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Delete author failed : %s", e.Error()),
			500,
		)
	}
	return
}

func (srv *AuthorsService) hashPassword(password string) (hashResult []byte, err error) {
	pwd := []byte(password)
	hashResult, err = bcrypt.GenerateFromPassword(pwd, 10)
	return
}
