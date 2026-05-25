package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"expense-tracker-api/database"
)

type DashboardData struct {
	TotalIncome       float64 `json:"totalIncome"`
	TotalExpense      float64 `json:"totalExpense"`
	NetBalance        float64 `json:"netBalance"`
	TotalTransactions int     `json:"totalTransactions"`
}

func GetDashboardStats(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().
		Value("user_id").(int)

	query := `
	SELECT
		COALESCE(
			SUM(
				CASE
					WHEN type='income'
					THEN amount
					ELSE 0
				END
			), 0
		),
		COALESCE(
			SUM(
				CASE
					WHEN type='expense'
					THEN amount
					ELSE 0
				END
			), 0
		),
		COUNT(*)
	FROM transactions
	WHERE user_id=$1
	`

	var stats DashboardData

	err := database.DB.QueryRow(
		context.Background(),
		query,
		userID,
	).Scan(
		&stats.TotalIncome,
		&stats.TotalExpense,
		&stats.TotalTransactions,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to fetch dashboard stats",
			http.StatusInternalServerError,
		)

		return
	}

	stats.NetBalance =
		stats.TotalIncome -
			stats.TotalExpense

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(stats)
}
