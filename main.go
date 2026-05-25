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
	"github.com/rs/cors"
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

	r.Handle(
		"/dashboard",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.GetDashboardStats,
			),
		),
	).Methods("GET")

	r.Handle(
		"/analytics/categories",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.GetCategoryAnalytics,
			),
		),
	).Methods("GET")

	r.Handle(
		"/analytics/trends",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.GetMonthlyTrends,
			),
		),
	).Methods("GET")

	r.Handle(
		"/transactions",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.CreateTransaction,
			),
		),
	).Methods("POST")

	r.Handle(
		"/transactions",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.GetTransactions,
			),
		),
	).Methods("GET")

	r.Handle(
		"/transactions/{id}",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.DeleteTransaction,
			),
		),
	).Methods("DELETE")

	r.Handle(
		"/accounts",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.CreateAccount,
			),
		),
	).Methods("POST")

	r.Handle(
		"/accounts",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.GetAccounts,
			),
		),
	).Methods("GET")

	r.Handle(
		"/categories",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.CreateCategory,
			),
		),
	).Methods("POST")

	r.Handle(
		"/categories",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.GetCategories,
			),
		),
	).Methods("GET")

	r.Handle(
		"/profile",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.GetProfile,
			),
		),
	).Methods("GET")

	r.Handle(
		"/profile",
		middleware.JWTAuth(
			http.HandlerFunc(
				handlers.UpdateProfile,
			),
		),
	).Methods("PUT")

	port := os.Getenv("PORT")

	log.Println("Server running on port", port)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"https://expensetrackergo.netlify.app",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},
		AllowedHeaders: []string{
			"*",
		},
	})

	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
