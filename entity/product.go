package entity

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MainImg     string    `json:"mainImg"`
	Price       int       `json:"price"`
	CategoryId  int       `json:"categoryId"`
	Category    string    `json:"category"`
	CreatedBy   int       `json:"createdBy,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}
