package handler

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/test_kompas/news_app/internal/pkg"
	"github.com/test_kompas/news_app/pkg/entity"
)

type ArticleHandler struct {
	authenticationSrv entity.AuthenticationService
	articleSrv        entity.ArticleService
}

func NewArticleHandler(authenticationSrv entity.AuthenticationService, articleSrv entity.ArticleService) *ArticleHandler {
	return &ArticleHandler{authenticationSrv, articleSrv}
}

func (handler *ArticleHandler) Find(c *fiber.Ctx) error {
	var err *pkg.Errors
	var e error

	authToken := c.Get("Authorization")
	if _, err := handler.authenticationSrv.Authorize(authToken); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	result := make([]entity.Article, 0)

	query := make(map[string]interface{})
	id := c.Query("id")
	name := c.Query("name")

	if id != "" {
		query["id"] = id
	}
	if name != "" {
		query["name"] = name
	}

	limit := c.Query("limit_per_page")
	page := c.Query("page_no")

	limitPerPage := 0
	pageNo := 0
	if limit != "" {
		limitPerPage, e = strconv.Atoi(c.Query("limit_per_page"))
		if e != nil {
			response := fiber.Map{
				"result": nil,
				"error":  "Invalid Parameter",
			}
			c.JSON(response)
			return c.SendStatus(400)
		}
	}

	if page != "" {
		pageNo, e = strconv.Atoi(c.Query("page_no"))
		if e != nil {
			response := fiber.Map{
				"result": nil,
				"error":  "Invalid Parameter",
			}
			c.JSON(response)
			return c.SendStatus(400)
		}
	}

	pagination := &entity.Pagination{
		Limit:  limitPerPage,
		PageNo: pageNo,
	}
	if result, err = handler.articleSrv.FindArticles(query, pagination); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	response := fiber.Map{
		"result": result,
		"error":  nil,
	}
	return c.JSON(response)
}

func (handler *ArticleHandler) AddArticle(c *fiber.Ctx) error {

	authToken := c.Get("Authorization")
	payload, err := handler.authenticationSrv.Authorize(authToken)
	if err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	body := entity.Article{}

	if err := json.Unmarshal(c.Body(), &body); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(500)
	}

	var authorId uint
	if data, ok := payload.Get("author_id"); ok {
		authorId = data.(uint)
	}

	if err := handler.articleSrv.AddArticle(&body, authorId); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	response := fiber.Map{
		"result": "success",
		"error":  nil,
	}
	return c.JSON(response)
}
