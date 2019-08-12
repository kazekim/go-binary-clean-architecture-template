package usecase

import (
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/logic"
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/service"
	"github.com/kazekim/go-binary-clean-architecture-template/model"
)

type GreetInput struct {
	Name string
	Prefix *string
}

type exampleUseCase struct {
	service *service.ExampleService
}

func NewExampleUseCase(service *service.ExampleService) *exampleUseCase {
	return &exampleUseCase{
		service: service,
	}
}

func (u exampleUseCase) Greet(input GreetInput) (*model.ExampleModel, error) {

	e, err := u.service.DoSomething(input.Name);
	if err != nil {
		return nil, err
	}

	year := 1900
	if e.YearOfBirth != nil {
		year = *e.YearOfBirth
	}
	age := logic.CalculateAge(year)

	m := model.ExampleModel{
		Name: e.Name,
		Age: &age,
	}

	return &m, nil
}
