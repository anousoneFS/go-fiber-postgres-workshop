package province

import (
	"errors"
	"html"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ProvinceRequest struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

func (p *ProvinceRequest) EscapeWhiteSpace() {
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.NameEn = html.EscapeString(strings.TrimSpace(p.NameEn))
}

func (p ProvinceRequest) Validate() error {
	if p.Name == "" || p.NameEn == "" {
		return errors.New("validate failed")
	}
	return nil
}

type handler struct {
	uc Usecase
}

func NewHandler(uc Usecase, app *fiber.App) {
	h := &handler{uc: uc}
	api := app.Group("/api/v1")
	api.Get("/provinces", h.GetAll)
	api.Post("/provinces", h.Create)
}

func (h handler) Create(c *fiber.Ctx) error {
	var body ProvinceRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	body.EscapeWhiteSpace()
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	if err := h.uc.Create(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpected error"})
	}
	// todo: get province
	return c.Status(fiber.StatusOK).JSON(body)
}

func (h handler) GetAll(c *fiber.Ctx) error {
	i, err := h.uc.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpected error"})
	}
	return c.Status(fiber.StatusOK).JSON(i)
}
