package main

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/province", GetProvince)
	app.Post("/province", CreateProvince)
	app.Listen(":3000")
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

type ProvinceResponse struct {
	Name string `json:"name"`
}

func GetProvince(c *fiber.Ctx) error {
	return c.SendString("hello makerbox")
}

func CreateProvince(c *fiber.Ctx) error {
	var p ProvinceRequest
	// map
	if err := c.BodyParser(&p); err != nil {
		// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
		return c.Status(fiber.StatusBadRequest).JSON("invalid body")
	}
	// validate
	if err := p.Validate(); err != nil {
		// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid"})
		return c.Status(fiber.StatusBadRequest).JSON("invalid body")
	}
	// insert into table province
	fmt.Printf("body: %+v\n", p)
	return c.SendString("Create success")
}
