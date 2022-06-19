package province

import "fmt"

type Usecase interface {
	GetAll() (p []Province, err error)
	Create(p ProvinceRequest) (id uint, err error)
	Update()
	GetByID()
}

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{repo: repo}
}

func (u usecase) Create(p ProvinceRequest) (id uint, err error) {
	province := Province{Name: p.Name, NameEn: p.NameEn}
	id, err = u.repo.Create(&province)
	if err != nil {
		fmt.Printf("usecase.GetAll():%v\n", err)
		return 0, err
	}
	return
}

func (u usecase) GetAll() (p []Province, err error) {
	i, err := u.repo.GetAll()
	if err != nil {
		// log
		fmt.Printf("usecase.GetAll():%v\n", err)
		return []Province{}, err
	}
	return i, err
}

// GetByID implements Usecase
func (*usecase) GetByID() {
	panic("unimplemented")
}

// Update implements Usecase
func (*usecase) Update() {
	panic("unimplemented")
}
