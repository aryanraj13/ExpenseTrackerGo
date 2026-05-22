package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"expense-tracker-api/database"
	"expense-tracker-api/models"

	"expense-tracker-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	query := `
	INSERT INTO users (name, email, password)
	VALUES ($1, $2, $3)
	`

	_, err = database.DB.Exec(
		context.Background(),
		query,
		user.Name,
		user.Email,
		string(hashedPassword),
	)

	if err != nil {
		http.Error(w, "User already exists or DB error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	var dbUser models.User

	query := `
	SELECT id, name, email, password
	FROM users
	WHERE email=$1
	`

	err = database.DB.QueryRow(
		context.Background(),
		query,
		user.Email,
	).Scan(
		&dbUser.ID,
		&dbUser.Name,
		&dbUser.Email,
		&dbUser.Password,
	)

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(dbUser.Password),
		[]byte(user.Password),
	)

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(dbUser.ID)

	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
