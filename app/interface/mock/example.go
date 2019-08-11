package mock

import "github.com/kazekim/go-binary-clean-architecture-template/app/domain/entity"

type exampleRepository struct {

}

func NewExampleRepository() *exampleRepository {
	return &exampleRepository{
	}
}

func (repo *exampleRepository) Greet(model entity.Example) (*string, error) {

	text := "Hello! " + model.Name + "!"
	return &text, nil
}