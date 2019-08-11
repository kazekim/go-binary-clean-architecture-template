package usecase

import (
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/entity"
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/logic"
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/service"
)

type exampleUseCase struct {
	service service.ExampleService
}

func NewExampleUseCase(service service.ExampleService) *exampleUseCase {
	return &exampleUseCase{
		service: service,
	}
}

func (useCase exampleUseCase) Greet(model entity.Example) (*string, error) {

	year := 1900
	if model.YearOfBirth != nil {
		year = *model.YearOfBirth
	}
	age := logic.CalculateAge(year)

	model.Age = &age

	return useCase.service.DoSomething(model)
}
