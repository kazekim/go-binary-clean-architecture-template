package service

import (
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/entity"
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/repository"
)

type ExampleService struct {
	repo repository.ExampleRepository
}

func NewExampleService(repo repository.ExampleRepository) *ExampleService {
	return &ExampleService{
		repo: repo,
	}
}

func (service *ExampleService) DoSomething(name string) (*entity.Example, error) {


	e, err := service.repo.Greet(name)
	if err != nil {
		return nil, err
	}

	return e, nil
}