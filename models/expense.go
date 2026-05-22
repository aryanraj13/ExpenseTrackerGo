package models

import "time"

type Expense struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Amount    float64 `json:"amount"`
	Category  string  `json:"category"`
	CreatedAt time.Time  `json:"created_at"`
	UserID    int     `json:"user_id"`
}
