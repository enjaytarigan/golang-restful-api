package entity

import "time"

type Category struct {
	ID int
	Name string
	CreatedAt time.Time
	CreatedBy int
}