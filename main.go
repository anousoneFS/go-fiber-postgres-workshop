package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/anousoneFS/go-workshop/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func failOnError(err error, msg string) {
	if err != nil {
		os.Exit(1)
	}
}

type SqlLogger struct {
	logger.Interface
}

func (s SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n================\n", sql)
}

var db *gorm.DB

func main() {
	cfg, err := config.LoadConfig("./")
	failOnError(err, "failed to load config")
	// gorm
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=Asia/Vientiane", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	dial := postgres.Open(dsn)
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}
	// if err = db.AutoMigrate(Gender{}, Test{}, Customer{}); err != nil {
	// 	fmt.Println(err)
	// }

	// if err = db.Migrator().CreateTable(&Gender{}); err != nil {
	// 	fmt.Println(err)
	// }
	// CreateGender("Male")
	// CreateGender("Female")
	// CreateGender("Other")
	// GetGender(1)
	// GetAllGender()
	// UpdateGender(1, "")
	// DeleteGender(1)

	// CreateTest("test1")
	// CreateTest("test2")
	// CreateTest("test3")
	// DeleteTest(2)
	// GetTest(1)
	// db.Migrator().CreateTable(&Customer{})

	// CreateCustomer("someone", 3)
	// GetCustomer(1)
	GetAllCustomer()
}

func GetCustomer(id uint) {
	customer := Customer{}
	tx := db.Preload(clause.Associations).Find(&customer, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Printf("customer: %+v\n", customer)
}

func GetAllCustomer() {
	customer := []Customer{}
	tx := db.Preload("Gender").Find(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _, i := range customer {
		fmt.Printf("%v|%v|%v\n", i.ID, i.Name, i.Gender.Name)
	}
}

func CreateCustomer(name string, genderId uint) {
	customer := Customer{Name: name, GenderID: genderId}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Printf("customer: %+v\n", customer)
}

type Customer struct {
	ID       uint
	Name     string
	Gender   Gender
	GenderID uint
}

func DeleteGender(id int) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Printf("gender:%+v\n", gender)
}

func GetGender(code int) {
	gender := Gender{}
	tx := db.Where("code=?", code).Find(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Printf("gender:%+v\n", gender)
}

func GetAllGender() {
	gender := []Gender{}
	tx := db.Find(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Printf("gender:%+v\n", gender)
}

func UpdateGender(code int, name string) {
	tx := db.Model(&Gender{}).Where("code = ?", code).Update("myname", name)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(code)
}

type Gender struct {
	Code int    `gorm:"primarykey"`
	Name string `gorm:"column:myname;size:8;not null"`
}

func CreateTest(name string) {
	test := Test{Name: name}
	db.Create(&test)
}

func GetTest(id int) {
	i := []Test{}
	db.Find(&i)
	fmt.Println(i)
}

func DeleteTest(id int) {
	db.Unscoped().Delete(&Test{}, id)
}

type Test struct {
	gorm.Model
	Name string
}
