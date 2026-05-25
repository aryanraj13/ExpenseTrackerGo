package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"expense-tracker-api/database"
	"expense-tracker-api/models"
)

func CreateCategory(
	w http.ResponseWriter,
	r *http.Request,
) {

	var category models.Category

	err := json.NewDecoder(
		r.Body,
	).Decode(&category)

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
	INSERT INTO categories (
		name,
		color,
		icon,
		user_id
	)
	VALUES ($1, $2, $3, $4)
	`

	_, err = database.DB.Exec(
		context.Background(),
		query,
		category.Name,
		category.Color,
		category.Icon,
		userID,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to create category",
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Category created",
		},
	)
}

func GetCategories(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().
		Value("user_id").(int)

	query := `
	SELECT
		id,
		name,
		color,
		icon,
		user_id,
		created_at
	FROM categories
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
			"Failed to fetch categories",
			http.StatusInternalServerError,
		)

		return
	}

	defer rows.Close()

	var categories []models.Category

	for rows.Next() {

		var category models.Category

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Color,
			&category.Icon,
			&category.UserID,
			&category.CreatedAt,
		)

		if err != nil {

			continue
		}

		categories = append(
			categories,
			category,
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		categories,
	)
}
