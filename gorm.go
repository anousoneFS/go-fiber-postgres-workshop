package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Province struct {
	gorm.Model
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

func Create(p ProvinceRequest) error {
	u := Province{Name: p.Name, NameEn: p.NameEn}
	tx := DB.Create(&u)
	if err := tx.Error; err != nil {
		fmt.Printf("CreateProvince():%v\n", err)
		return err
	}
	fmt.Printf("result: %+v\n", u.ID)
	return nil
}
