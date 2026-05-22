package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"expense-tracker-api/database"
	"expense-tracker-api/models"

	"github.com/gorilla/mux"
)

func CreateExpense(w http.ResponseWriter, r *http.Request) {

	var expense models.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(int)

	query := `
	INSERT INTO expenses (title, amount, category, user_id)
	VALUES ($1, $2, $3, $4)
	`

	_, err = database.DB.Exec(
		context.Background(),
		query,
		expense.Title,
		expense.Amount,
		expense.Category,
		userID,
	)

	if err != nil {
		http.Error(w, "Failed to create expense", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Expense created successfully",
	})
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id").(int)

	query := `
	SELECT id, title, amount, category, created_at, user_id
	FROM expenses
	WHERE user_id=$1
	ORDER BY created_at DESC
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {
		http.Error(w, "Failed to fetch expenses", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {

		var expense models.Expense

		err := rows.Scan(
			&expense.ID,
			&expense.Title,
			&expense.Amount,
			&expense.Category,
			&expense.CreatedAt,
			&expense.UserID,
		)

		if err != nil {
			http.Error(w, "Error scanning expenses", http.StatusInternalServerError)
			return
		}

		expenses = append(expenses, expense)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(expenses)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {

	var expense models.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(int)

	params := mux.Vars(r)

	expenseID := params["id"]

	query := `
	UPDATE expenses
	SET title=$1, amount=$2, category=$3
	WHERE id=$4 AND user_id=$5
	`

	_, err = database.DB.Exec(
		context.Background(),
		query,
		expense.Title,
		expense.Amount,
		expense.Category,
		expenseID,
		userID,
	)

	if err != nil {
		http.Error(w, "Failed to update expense", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Expense updated successfully",
	})
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id").(int)

	params := mux.Vars(r)

	expenseID := params["id"]

	query := `
	DELETE FROM expenses
	WHERE id=$1 AND user_id=$2
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		expenseID,
		userID,
	)

	if err != nil {
		http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Expense deleted successfully",
	})
}

func GetExpenseSummary(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id").(int)

	query := `
	SELECT category, COALESCE(SUM(amount), 0)
	FROM expenses
	WHERE user_id=$1
	GROUP BY category
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {
		http.Error(w, "Failed to fetch summary", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	type Summary struct {
		Category string  `json:"category"`
		Total    float64 `json:"total"`
	}

	var summary []Summary

	for rows.Next() {

		var item Summary

		err := rows.Scan(
			&item.Category,
			&item.Total,
		)

		if err != nil {
			http.Error(w, "Error scanning summary", http.StatusInternalServerError)
			return
		}

		summary = append(summary, item)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(summary)
}
