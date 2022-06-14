package province

import "fmt"

type usecase struct {
	repo Repository
}

type Usecase interface {
	Create(p ProvinceRequest) error
	GetAll() ([]Province, error)
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{repo: repo}
}

func (u usecase) Create(body ProvinceRequest) error {
	province := Province{Name: body.Name, NameEn: body.NameEn}
	if err := u.repo.Create(&province); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u usecase) GetAll() (i []Province, err error) {
	i, err = u.repo.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	return
}
