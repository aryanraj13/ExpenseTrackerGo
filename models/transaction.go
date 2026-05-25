package models

import "time"

type Transaction struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
