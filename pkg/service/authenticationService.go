package service

import (
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/test_kompas/news_app/configs"
	"github.com/test_kompas/news_app/internal/pkg"
	"github.com/test_kompas/news_app/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService struct {
	authRepo   entity.AuthenticationRepository
	authorRepo entity.AuthorsRepository
}

func NewAuthenticationService(authRepo entity.AuthenticationRepository, authorRepo entity.AuthorsRepository) *AuthenticationService {
	return &AuthenticationService{authRepo, authorRepo}
}

func (srv *AuthenticationService) AuthorLogin(author *entity.Authors) (jwtToken string, err *pkg.Errors) {

	if author.Username == "" || author.Password == "" {
		err = pkg.NewError("username and password field is required", 400)
		return
	}
	getAuthor, e := srv.authorRepo.FindByUsername(author.Username)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Could not retrieve author data : %s", e.Error()),
			500,
		)
		return
	}
	if getAuthor.ID == 0 {
		err = pkg.NewError(
			fmt.Sprintf("Username \"%s\" does not exists", author.Username),
			404,
		)
		return
	}

	if e := srv.validatePassword(getAuthor.Password, author.Password); e != nil {
		err = pkg.NewError("Wrong password", 400)
		return
	}

	authentication := entity.Authentication{
		AuthorID:   getAuthor.ID,
		IssuedAt:   time.Now().Format("2006-01-02 15:04:05"),
		ExpiryTime: 36000,
	}

	jwtToken, e = srv.generateToken(
		map[string]interface{}{
			"iss":       getAuthor.Username,
			"sub":       getAuthor.Username,
			"exp":       authentication.ExpiryTime,
			"iat":       authentication.IssuedAt,
			"author_id": authentication.AuthorID,
			"iat_unix":  time.Now().Unix(),
		},
	)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Generate token failed : %s", e.Error()),
			500,
		)
		return
	}

	authentication.Token = jwtToken

	if e := srv.authRepo.AddToken(&authentication); e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Author Login failed : %s", e.Error()),
			500,
		)
	}
	return
}

func (srv *AuthenticationService) Authorize(jwtToken string) (result jwt.Token, err *pkg.Errors) {

	if jwtToken == "" {
		err = pkg.NewError("Access token is required", 400)
		return
	}

	getToken, e := srv.authRepo.FindBytoken(jwtToken)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Get token failed : %s", e.Error()),
			500,
		)
		return
	}
	if getToken.ID == 0 {
		err = pkg.NewError("Unauthorized, invalid access token", 401)
		return
	}

	result, e = srv.validateToken(jwtToken)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Parse jwt failed : %s", e.Error()),
			500,
		)
		return
	}

	issuedAt, e := time.Parse("2006-01-02 15:04:05", getToken.IssuedAt)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Parse time failed : %s", e.Error()),
			500,
		)
		return
	}

	timeNow := time.Now().Unix()
	expiryTime := issuedAt.Unix() + int64(getToken.ExpiryTime)
	if timeNow > expiryTime {
		err = pkg.NewError("Unauthorized, Access token expired", 401)
	}
	return
}

func (srv *AuthenticationService) validateToken(payload string) (token jwt.Token, err error) {

	key := []byte(configs.HASHKEY)

	jwkKey, err := jwk.New(key)
	if err != nil {
		return
	}

	token, err = jwt.Parse(
		[]byte(payload),
		jwt.WithVerify(jwa.HS256, jwkKey),
	)
	return
}

func (srv *AuthenticationService) generateToken(payload map[string]interface{}) (jwtToken string, err error) {
	jwtHeader := jws.NewHeaders()
	jwtHeader.Set("typ", "JWT")
	jwtHeader.Set("alg", jwa.HS256)

	key := []byte(configs.HASHKEY)

	jwkKey, err := jwk.New(key)
	if err != nil {
		err = pkg.NewError(err.Error(), 403)
		return
	}
	token := jwt.New()
	for k, v := range payload {
		token.Set(k, v)
	}
	tokenPayload, err := jwt.Sign(token, jwa.HS256, jwkKey, jwt.WithHeaders(jwtHeader))
	if err != nil {
		err = pkg.NewError(err.Error(), 403)
		return
	}
	jwtToken = string(tokenPayload)
	return
}

func (srv *AuthenticationService) validatePassword(hashedPassword, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return
}
