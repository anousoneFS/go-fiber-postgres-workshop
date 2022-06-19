package main

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

var DB *gorm.DB

func main() {
	dsn := "postgres://oiaaglbm:M7yp7cg1uAG4UpiVazViExpoYwnZTdIw@tiny.db.elephantsql.com/oiaaglbm"
	dial := postgres.Open(dsn)
	var err error
	DB, err = gorm.Open(dial)
	if err != nil {
		panic(err)
	}
	if err = DB.AutoMigrate(Province{}); err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Get("/province", GetProvince)
	app.Post("/province", CreateProvince)
	app.Listen(":3000")
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
	if err := Create(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Printf("body: %+v\n", p)
	return c.SendString("Create success")
}
