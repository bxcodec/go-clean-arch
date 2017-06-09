package models

import "time"

type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Tag       string    `json:"tag"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
