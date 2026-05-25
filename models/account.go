package models

import "time"

type Account struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Balance   float64   `json:"balance"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
