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

	// CreateUser("daky", "daky@gmail.com", 23)
	// GetAllUser()
	// GetUserByID(8)
	// GetUserByName("anousone")
	UpdateUser(8, "makerbox")
}

// table: User
type User struct {
	Name  string `gorm:"column:myname;size:8;not null"`
	Email string `gorm:"not null"`
	Age   int    `gorm:"not null"`
	gorm.Model
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
	fmt.Printf("result: %+v\n", u.ID)
}

func GetAllUser() {
	user := []User{}
	if err := db.Find(&user).Error; err != nil {
		fmt.Printf("GetAllUser():%v\n", err)
		return
	}
	for index, item := range user {
		fmt.Printf("%v name:%v, email:%v, age:%v\n", index, item.Name, item.Email, item.Age)
	}
	// fmt.Printf("user:%+v\n", user)
}

func GetUserByID(id int) {
	user := User{}
	tx := db.Where("id=?", id).Find(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Printf("name:%v, email:%v, age:%v\n", user.Name, user.Email, user.Age)
}

func GetUserByName(name string) {
	user := User{}
	tx := db.Where("myname=?", name).Find(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Printf("name:%v, email:%v, age:%v\n", user.Name, user.Email, user.Age)
}

func UpdateUser(id int, name string) {
	tx := db.Model(&User{}).Where("id = ?", id).Update("myname", name)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Printf("update success")
}
