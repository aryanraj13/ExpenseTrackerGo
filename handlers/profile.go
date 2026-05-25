package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"expense-tracker-api/database"
)

type Profile struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetProfile(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().
		Value("user_id").(int)

	query := `
	SELECT
		id,
		name,
		email
	FROM users
	WHERE id=$1
	`

	var profile Profile

	err := database.DB.QueryRow(
		context.Background(),
		query,
		userID,
	).Scan(
		&profile.ID,
		&profile.Name,
		&profile.Email,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to fetch profile",
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		profile,
	)
}

func UpdateProfile(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().
		Value("user_id").(int)

	var profile Profile

	err := json.NewDecoder(
		r.Body,
	).Decode(&profile)

	if err != nil {

		http.Error(
			w,
			"Invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	query := `
	UPDATE users
	SET name=$1,
	email=$2
	WHERE id=$3
	`

	_, err = database.DB.Exec(
		context.Background(),
		query,
		profile.Name,
		profile.Email,
		userID,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to update profile",
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Profile updated",
		},
	)
}
