package handler

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/test_kompas/news_app/internal/pkg"
	"github.com/test_kompas/news_app/pkg/entity"
)

type AuthorsHandler struct {
	authenticationSrv entity.AuthenticationService
	authorSrv         entity.AuthorsService
}

func NewAuthorsHandler(authenticationSrv entity.AuthenticationService, authorSrv entity.AuthorsService) *AuthorsHandler {
	return &AuthorsHandler{authenticationSrv, authorSrv}
}

func (handler *AuthorsHandler) Find(c *fiber.Ctx) error {
	var err *pkg.Errors

	authToken := c.Get("Authorization")
	if _, err := handler.authenticationSrv.Authorize(authToken); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	result := make([]entity.Authors, 0)

	query := make(map[string]interface{})
	id := c.Query("id")
	name := c.Query("name")

	if id != "" {
		query["id"] = id
	}
	if name != "" {
		query["name"] = name
	}

	limitPerPage, e := strconv.Atoi(c.Query("limit_per_page"))
	if e != nil {
		response := fiber.Map{
			"result": nil,
			"error":  "Invalid Parameter",
		}
		c.JSON(response)
		return c.SendStatus(400)
	}
	pageNo, e := strconv.Atoi(c.Query("page_no"))
	if e != nil {
		response := fiber.Map{
			"result": nil,
			"error":  "Invalid Parameter",
		}
		c.JSON(response)
		return c.SendStatus(400)
	}

	pagination := &entity.Pagination{
		Limit:  limitPerPage,
		PageNo: pageNo,
	}
	if result, err = handler.authorSrv.FindAuthors(query, pagination); err != nil {
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

func (handler *AuthorsHandler) AddAuthor(c *fiber.Ctx) error {

	body := entity.Authors{}

	if err := json.Unmarshal(c.Body(), &body); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(500)
	}

	if _, err := handler.authorSrv.AddAuthor(&body); err != nil {
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

func (handler *AuthorsHandler) DeleteAuthor(c *fiber.Ctx) error {

	authToken := c.Get("Authorization")
	if _, err := handler.authenticationSrv.Authorize(authToken); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	id, e := strconv.Atoi(c.Query("id"))
	if e != nil {
		response := fiber.Map{
			"result": nil,
			"error":  "Invalid Parameter",
		}
		c.JSON(response)
		return c.SendStatus(400)
	}

	if err := handler.authorSrv.DeleteAuthor(uint(id)); err != nil {
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
