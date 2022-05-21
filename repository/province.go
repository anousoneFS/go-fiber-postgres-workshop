package repository

import (
	"gorm.io/gorm"
)

type ProvinceDB struct {
	gorm.Model
	Name   string `gorm:"type:text;not null"`
	NameEn string `gorm:"type:text"`
}

type ProvinceRepository struct {
	DB *gorm.DB
}

func (a *ProvinceRepository) Create(p ProvinceDB) error {
	if err := a.DB.Create(&p).Error; err != nil {
		return err
	}
	return nil
}
