package entity

import "time"

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
