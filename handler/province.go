package handler

import (
	"net/http"

	"github.com/anousoneFS/administrative-divisions/service"
	"github.com/gofiber/fiber/v2"
)

type ProvinceHandler struct {
	Svc *service.ProvinceService
}

func (p *ProvinceHandler) Add(c *fiber.Ctx) error {
	var province service.Province
	if err := c.BodyParser(&province); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": ErrBodyParserFailure})
	}

	err := p.Svc.Add(province)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"error": err})
	}
	return c.Status(http.StatusCreated).JSON(&fiber.Map{"result": Created})
}

func (p *ProvinceHandler) SetupRoutes(app *fiber.App) {
	app.Post("/api/province-add", p.Add)
}
