package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"expense-tracker-api/database"
	"expense-tracker-api/models"

	"github.com/gorilla/mux"
)

func CreateTransaction(
	w http.ResponseWriter,
	r *http.Request,
) {

	var transaction models.Transaction

	err := json.NewDecoder(
		r.Body,
	).Decode(&transaction)

	if err != nil {

		http.Error(
			w,
			"Invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	userID := r.Context().
		Value("user_id").(int)

	query := `
	INSERT INTO transactions (
		title,
		type,
		amount,
		category,
		description,
		user_id
	)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = database.DB.Exec(
		context.Background(),
		query,
		transaction.Title,
		transaction.Type,
		transaction.Amount,
		transaction.Category,
		transaction.Description,
		userID,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to create transaction",
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Transaction created",
		},
	)
}

func GetTransactions(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().
		Value("user_id").(int)

	query := `
	SELECT
		id,
		title,
		type,
		amount,
		category,
		description,
		user_id,
		created_at
	FROM transactions
	WHERE user_id=$1
	ORDER BY created_at DESC
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to fetch transactions",
			http.StatusInternalServerError,
		)

		return
	}

	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {

		var transaction models.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.Title,
			&transaction.Type,
			&transaction.Amount,
			&transaction.Category,
			&transaction.Description,
			&transaction.UserID,
			&transaction.CreatedAt,
		)

		if err != nil {

			continue
		}

		transactions = append(
			transactions,
			transaction,
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		transactions,
	)
}

func DeleteTransaction(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().
		Value("user_id").(int)

	id := mux.Vars(r)["id"]

	query := `
	DELETE FROM transactions
	WHERE id=$1
	AND user_id=$2
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		id,
		userID,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to delete transaction",
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Transaction deleted",
		},
	)
}
