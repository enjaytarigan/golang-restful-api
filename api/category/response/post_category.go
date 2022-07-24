package response

import "brodo-demo/entity"

type PostCategoryResponse struct {
	Category *entity.Category `json:"category"`
}