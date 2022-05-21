package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/anousoneFS/administrative-divisions/handler"
	"github.com/anousoneFS/administrative-divisions/migration"
	"github.com/anousoneFS/administrative-divisions/repository"
	"github.com/anousoneFS/administrative-divisions/service"
)

var (
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
	DSN     = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Vientiane", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = migration.MigrateDB(db)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	provinceRepo := &repository.ProvinceRepository{DB: db}
	provinceSvc := &service.ProvinceService{Repo: provinceRepo}
	provinceHandler := &handler.ProvinceHandler{Svc: provinceSvc}

	provinceHandler.SetupRoutes(app)

	app.Listen(":8080")
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}
	return db, nil
}
