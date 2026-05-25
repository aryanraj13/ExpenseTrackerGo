package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"expense-tracker-api/database"
	"expense-tracker-api/models"
)

func CreateAccount(
	w http.ResponseWriter,
	r *http.Request,
) {

	var account models.Account

	err := json.NewDecoder(
		r.Body,
	).Decode(&account)

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
	INSERT INTO accounts (
		name,
		type,
		balance,
		user_id
	)
	VALUES ($1, $2, $3, $4)
	`

	_, err = database.DB.Exec(
		context.Background(),
		query,
		account.Name,
		account.Type,
		account.Balance,
		userID,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to create account",
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Account created",
		},
	)
}

func GetAccounts(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().
		Value("user_id").(int)

	query := `
	SELECT
		id,
		name,
		type,
		balance,
		user_id,
		created_at
	FROM accounts
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
			"Failed to fetch accounts",
			http.StatusInternalServerError,
		)

		return
	}

	defer rows.Close()

	var accounts []models.Account

	for rows.Next() {

		var account models.Account

		err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.Type,
			&account.Balance,
			&account.UserID,
			&account.CreatedAt,
		)

		if err != nil {

			continue
		}

		accounts = append(
			accounts,
			account,
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		accounts,
	)
}
