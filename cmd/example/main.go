package main

import (
	"fmt"
	"github.com/kazekim/go-binary-clean-architecture-template/app/domain/service"
	"github.com/kazekim/go-binary-clean-architecture-template/app/interface/mock"
	"github.com/kazekim/go-binary-clean-architecture-template/app/usecase"
)

func main() {

	repo := mock.NewExampleRepository()
	exampleService := service.NewExampleService(repo)
	useCase := usecase.NewExampleUseCase(exampleService)

	example := usecase.GreetInput{
		Name:        "Kim",
	}
	m, err := useCase.Greet(example)
	if err != nil {
		panic(err)
	}

	fmt.Println("Output : ", m.Name)
}
