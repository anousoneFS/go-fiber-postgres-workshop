package main

import (
	"fmt"

	"github.com/anousoneFS/go-fiber-postgres-workshop/province"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
	if err = DB.AutoMigrate(province.Province{}); err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Get("/province", GetAllProvince)
	app.Post("/province", CreateProvince)
	app.Listen(":3000")
}

func GetAllProvince(c *fiber.Ctx) error {
	repo := province.NewRepository(DB)
	p, err := repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	// map to province response model
	return c.Status(fiber.StatusOK).JSON(p)
}

func GetProvince(c *fiber.Ctx) error {
	return c.SendString("hello makerbox")
}

func CreateProvince(c *fiber.Ctx) error {
	repo := province.NewRepository(DB)
	var p province.ProvinceRequest
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
	if err := repo.Create(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Printf("body: %+v\n", p)
	return c.SendString("Create success")
}
