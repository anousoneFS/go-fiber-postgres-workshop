package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://oiaaglbm:M7yp7cg1uAG4UpiVazViExpoYwnZTdIw@tiny.db.elephantsql.com/oiaaglbm"
	dial := postgres.Open(dsn)
	db, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(User{}); err != nil {
		panic(err)
	}

	// db.Migrator().CreateTable(User{})
}

// table: User
type User struct {
	Name  string
	Email string
	Age   int
}
