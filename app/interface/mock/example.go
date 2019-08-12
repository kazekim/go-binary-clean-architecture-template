package mock

import "github.com/kazekim/go-binary-clean-architecture-template/app/domain/entity"

type exampleRepository struct {

}

func NewExampleRepository() *exampleRepository {
	return &exampleRepository{
	}
}

func (repo *exampleRepository) Greet(name string) (*entity.Example, error) {

	text := "Hello! " + name + "!"
	yob := 1986

	e := entity.Example{
		Name: text,
		YearOfBirth: &yob,
	}
	return &e, nil
}