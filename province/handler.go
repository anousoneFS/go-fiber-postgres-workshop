package province

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	us Usecase
}

func NewHandler(app *fiber.App, usecase Usecase) {
	h := &handler{
		us: usecase,
	}

	app.Get("/province", h.GetAll)
	app.Post("/province", h.Create)
}

type ProvinceRequest struct {
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

func (p ProvinceRequest) Validate() error {
	if p.NameEn == "" || p.Name == "" {
		return errors.New("invalid name_en")
	}
	return nil
}

func (h handler) GetAll(c *fiber.Ctx) error {
	response, err := h.us.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h handler) Create(c *fiber.Ctx) error {
	var body ProvinceRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid body")
	}
	// validate
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid body")
	}
	// create
	id, err := h.us.Create(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successfull.",
		"id":      id,
	})
}
