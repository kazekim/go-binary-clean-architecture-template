package repository

import "github.com/kazekim/go-binary-clean-architecture-template/app/domain/entity"

type ExampleRepository interface {
	 Greet(model entity.Example) (*string, error)
}