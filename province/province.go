package province

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ProvinceRequest struct {
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

func (p ProvinceRequest) Validate() error {
	if p.NameEn == "" || p.Name == "" {
		return errors.New("invalid name_en")
	}
	return nil
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

type Province struct {
	gorm.Model
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

func (r repository) GetAll() (p []Province, err error) {
	if err = r.db.Find(&p).Error; err != nil {
		fmt.Printf("GetAllProvince():%v\n", err)
		return
	}
	return p, nil
}

func (r repository) Create(p ProvinceRequest) error {
	u := Province{Name: p.Name, NameEn: p.NameEn}
	tx := r.db.Create(&u)
	if err := tx.Error; err != nil {
		fmt.Printf("CreateProvince():%v\n", err)
		return err
	}
	fmt.Printf("result: %+v\n", u.ID)
	return nil
}
