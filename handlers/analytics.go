package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"expense-tracker-api/database"
)

type CategoryAnalytics struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
}

func GetCategoryAnalytics(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().
		Value("user_id").(int)

	query := `
	SELECT
		category,
		SUM(amount)
	FROM transactions
	WHERE user_id=$1
	AND type='expense'
	GROUP BY category
	ORDER BY SUM(amount) DESC
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to fetch analytics",
			http.StatusInternalServerError,
		)

		return
	}

	defer rows.Close()

	var analytics []CategoryAnalytics

	for rows.Next() {

		var item CategoryAnalytics

		err := rows.Scan(
			&item.Category,
			&item.Total,
		)

		if err != nil {

			continue
		}

		analytics = append(
			analytics,
			item,
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(analytics)
}

type MonthlyTrend struct {
	Month string  `json:"month"`
	Total float64 `json:"total"`
}

func GetMonthlyTrends(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().
		Value("user_id").(int)

	query := `
	SELECT
		TO_CHAR(created_at, 'Mon') as month,
		SUM(amount)
	FROM transactions
	WHERE user_id=$1
	AND type='expense'
	GROUP BY month
	ORDER BY MIN(created_at)
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to fetch trends",
			http.StatusInternalServerError,
		)

		return
	}

	defer rows.Close()

	var trends []MonthlyTrend

	for rows.Next() {

		var item MonthlyTrend

		err := rows.Scan(
			&item.Month,
			&item.Total,
		)

		if err != nil {

			continue
		}

		trends = append(
			trends,
			item,
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(trends)
}
