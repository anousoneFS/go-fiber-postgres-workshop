package repository

import (
	"gorm.io/gorm"
)

type DistrictDB struct {
	gorm.Model
	Name       string `gorm:"type:text;not null"`
	NameEn     string `gorm:"type:text"`
	ProvinceID uint
	Province   ProvinceDB
}

type DistrictRepository struct {
	DB *gorm.DB
}
