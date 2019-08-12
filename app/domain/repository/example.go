package repository

import "github.com/kazekim/go-binary-clean-architecture-template/app/domain/entity"

type ExampleRepository interface {
	Greet(name string) (*entity.Example, error)
}