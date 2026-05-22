package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"expense-tracker-api/database"
	"expense-tracker-api/handlers"

	"expense-tracker-api/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}

	database.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Expense Tracker API Running")
	})

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	r.Handle(
		"/expenses",
		middleware.JWTAuth(http.HandlerFunc(handlers.CreateExpense)),
	).Methods("POST")

	r.Handle(
		"/expenses",
		middleware.JWTAuth(http.HandlerFunc(handlers.GetExpenses)),
	).Methods("GET")

	r.Handle(
		"/expenses/{id}",
		middleware.JWTAuth(http.HandlerFunc(handlers.UpdateExpense)),
	).Methods("PUT")

	r.Handle(
		"/expenses/{id}",
		middleware.JWTAuth(http.HandlerFunc(handlers.DeleteExpense)),
	).Methods("DELETE")

	r.Handle(
		"/expenses/summary",
		middleware.JWTAuth(http.HandlerFunc(handlers.GetExpenseSummary)),
	).Methods("GET")

	port := os.Getenv("PORT")

	log.Println("Server running on port", port)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
