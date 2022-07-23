package common

type SuccessResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Status: true,
		Data:   data,
	}
}
