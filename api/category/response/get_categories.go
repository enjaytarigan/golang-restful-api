package response

import "brodo-demo/entity"

type GetCategories struct {
	Categories []entity.Category `json:"categories"`
}