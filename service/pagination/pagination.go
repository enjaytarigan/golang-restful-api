package pagination

import (
	"math"
)

func CreatePagination(page, size int) (offset int, limit int, currentPage int) {
	currentPage = page
	limit = size
	
	if page <= 0 {
		// 1 is default page
		currentPage = 1
	}

	if size <= 0 {
		// 5 is default row per page
		limit = 5
	}

	offset = (currentPage - 1) * limit

	return offset, limit, currentPage
}

type Pagination struct {
	TotalItems  int `json:"totalItems,omitempty"`
	TotalPages  int `json:"totalPages,omitempty"`
	CurrentPage int `json:"currentPage,omitempty"`
}

func NewPagination(totalItems int, page int, size int) *Pagination {
	totalPages := math.Ceil(float64(totalItems) / float64(size))

	return &Pagination{
		TotalItems:  totalItems,
		TotalPages:  int(totalPages),
		CurrentPage: page,
	}
}
