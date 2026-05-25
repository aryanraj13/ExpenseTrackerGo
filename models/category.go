package models

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Icon      string    `json:"icon"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
