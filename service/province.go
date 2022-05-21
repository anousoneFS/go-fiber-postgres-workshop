package service

import (
	"github.com/anousoneFS/administrative-divisions/repository"
)

type Province struct {
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

type ProvinceService struct {
	Repo *repository.ProvinceRepository
}

func (p *ProvinceService) Add(province Province) error {
	if err := p.Repo.Create(repository.ProvinceDB{
		Name:   province.Name,
		NameEn: province.NameEn,
	}); err != nil {
		return err
	}
	return nil
}
