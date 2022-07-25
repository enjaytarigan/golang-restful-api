package response

import "brodo-demo/entity"

type ProductResponse struct {
	Product *entity.Product `json:"product"`
}
