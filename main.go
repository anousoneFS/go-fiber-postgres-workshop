package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "postgres://oiaaglbm:M7yp7cg1uAG4UpiVazViExpoYwnZTdIw@tiny.db.elephantsql.com/oiaaglbm"
	dial := postgres.Open(dsn)
	var err error
	db, err = gorm.Open(dial)
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(User{}); err != nil {
		panic(err)
	}
	// db.Migrator().CreateTable(User{})

	CreateUser("anousone", "myemail@gmail.com", 23)
}

// table: User
type User struct {
	ID    int    // default ID is primarykey
	Name  string `gorm:"column:myname;size:8;not null"`
	Email string `gorm:"not null"`
	Age   int    `gorm:"not null"`
}

func (u User) TableName() string {
	return "myuser"
}

func CreateUser(name, email string, age int) {
	u := User{Name: name, Email: email, Age: age}
	tx := db.Create(&u)
	if err := tx.Error; err != nil {
		fmt.Printf("CreateUser():%v\n", err)
	}
	fmt.Printf("result: %+v\n", u)
}
