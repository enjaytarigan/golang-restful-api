package request

type PostProductRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	CategoryId  int    `json:"categoryId" form:"categoryId"`
}
