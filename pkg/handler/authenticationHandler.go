package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/test_kompas/news_app/internal/pkg"
	"github.com/test_kompas/news_app/pkg/entity"
)

type AuthenticationHandler struct {
	authenticationSrv entity.AuthenticationService
}

func NewAuthenticationHandler(authenticationSrv entity.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{authenticationSrv}
}

func (handler *AuthenticationHandler) Login(c *fiber.Ctx) error {
	var err *pkg.Errors

	body := entity.Authors{}

	if err := json.Unmarshal(c.Body(), &body); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(500)
	}

	jwtToken, err := handler.authenticationSrv.AuthorLogin(&body)
	if err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	response := fiber.Map{
		"result": jwtToken,
		"error":  nil,
	}
	return c.JSON(response)
}
