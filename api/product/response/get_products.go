package response

import (
	"brodo-demo/entity"
	"brodo-demo/service/pagination"
)

type GetProductsResponse struct {
	Products []entity.Product `json:"products"`
	*pagination.Pagination
}