package models

type Expense struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Amount    float64 `json:"amount"`
	Category  string  `json:"category"`
	CreatedAt string  `json:"created_at"`
	UserID    int     `json:"user_id"`
}
