# Expense Tracker Backend API 🚀

A RESTful backend API for the Expense Tracker application built using **Golang**, **PostgreSQL**, and **JWT Authentication**. The API provides secure authentication and complete expense CRUD operations with protected routes and analytics support.

## 🌐 Live API

Backend URL: https://expensetrackergo.onrender.com

---

# ✨ Features

- 🔐 JWT Authentication
- 👤 User Registration & Login
- ➕ Create Expenses
- 📥 Fetch Expenses
- ✏️ Update Expenses
- 🗑 Delete Expenses
- 📊 Expense Summary Analytics
- 🛡 Protected Routes Middleware
- ⚡ Fast REST API using Gorilla Mux
- ☁️ Deployed on Render

---

# 🛠 Tech Stack

- Golang
- PostgreSQL
- Gorilla Mux
- JWT Authentication
- pgx PostgreSQL Driver
- Render Deployment

---

# 📂 Project Structure

```bash
expense-tracker-api/
│
├── database/
├── handlers/
├── middleware/
├── models/
├── routes/
├── .env
├── go.mod
└── main.go
```

---

# ⚙️ Environment Variables

Create a `.env` file:

```env id="pqv6w8"
PORT=8080

DATABASE_URL=your_postgresql_database_url

JWT_SECRET=your_secret_key
```

---

# 🧪 Running Locally

## Clone Repository

```bash id="ch5l9s"
git clone https://github.com/aryanraj13/ExpenseTrackerGo.git
```

---

## Install Dependencies

```bash id="fq21am"
go mod tidy
```

---

## Run Server

```bash id="v6y0kl"
go run main.go
```

Server runs on:

```bash id="gq4nbe"
http://localhost:8080
```

---

# 🔑 API Endpoints

## Authentication

| Method | Endpoint | Description |
|---|---|---|
| POST | `/register` | Register a new user |
| POST | `/login` | Login user |

---

## Expenses

| Method | Endpoint | Description |
|---|---|---|
| GET | `/expenses` | Get all expenses |
| POST | `/expenses` | Create expense |
| PUT | `/expenses/:id` | Update expense |
| DELETE | `/expenses/:id` | Delete expense |
| GET | `/expenses/summary` | Get expense analytics |

---

# 🔒 Authentication

Protected routes require JWT token:

```http id="mb1r0n"
Authorization: Bearer your_jwt_token
```

---

# ☁️ Deployment

Backend deployed on:

- Render

---

# 👨‍💻 Author

Aryan Rajput

GitHub: https://github.com/aryanraj13