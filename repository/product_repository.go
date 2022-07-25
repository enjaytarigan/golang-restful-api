package repository

import "brodo-demo/entity"

type FindAllProductsParam struct {
	Limit int
	Skip  int
}

type ProductRepository interface {
	InsertOne(product entity.Product) (*entity.Product, error)
	VerifyProductTypeIsExists(productTypeId int) error
	FindAllAndCount(params FindAllProductsParam) ([]entity.Product, int, error)
}
