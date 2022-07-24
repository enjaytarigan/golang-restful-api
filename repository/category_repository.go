package repository

import "brodo-demo/entity"

type CategoryRepository interface {
	InsertOne(category entity.Category)	(*entity.Category, error)
	UpdateById(category entity.Category) (*entity.Category, error)
	// FindAll() ([]entity.Category, error)
	FindById(Id int) (entity.Category, error)
}