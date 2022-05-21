package repository

import (
	"gorm.io/gorm"
)

type VillageDB struct {
	gorm.Model
	Name       string `gorm:"type:text;not null"`
	NameEn     string `gorm:"type:text"`
	DistrictID uint
	District   DistrictDB
}
