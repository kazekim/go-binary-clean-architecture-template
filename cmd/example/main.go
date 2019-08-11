package main

import (
	"fmt"
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/entity"
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/service"
	"github.com/kazekim/go-binary-clean-architecture-template/app/interface/mock"
	"github.com/kazekim/go-binary-clean-architecture-template/app/usecase"
)

func main() {

	repo := mock.NewExampleRepository()
	exampleService := service.NewExampleService(repo)
	useCase := usecase.NewExampleUseCase(exampleService)

	var yob int
	yob = 1986
	example := entity.Example{
		Name:        "Kim",
		YearOfBirth: &yob,
	}
	text, err := useCase.Greet(example)
	if err != nil {
		panic(err)
	}

	fmt.Println("Output : ", *text)
}
