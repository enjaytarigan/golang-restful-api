package repository

import "brodo-demo/entity"

type FindAllProductsParam struct {
	Limit int
	Skip  int
	MinPrice int
	MaxPrice int
}

type ProductRepository interface {
	InsertOne(product entity.Product) (*entity.Product, error)
	FindAllAndCount(params FindAllProductsParam) ([]entity.Product, int, error)
	FindById(productId int) (entity.Product, error)
}
