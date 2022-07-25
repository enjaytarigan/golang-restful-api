package entity

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MainImg     string    `json:"mainImg"`
	Price       int       `json:"price"`
	CategoryId  int       `json:"categoryId"`
	CreatedBy   int       `json:"createdBy"`
	Type        *int       `json:"type"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
}
