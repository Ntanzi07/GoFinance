# 💸 GoFinance — Financial Control API with Go + MariaDB

A practical challenge to build a **RESTful API for personal financial management**, developed in **Go (Golang)** with persistence in **MariaDB**.  
The goal is to create a modular and high-performance application to manage **income**, **expenses**, and **total balance**, applying best practices in architecture, testing, and documentation.

---

## 🎯 Objective

Implement an API that allows users to:
- Register **financial transactions** (income and expenses);
- Calculate **total balance**;
- Filter transactions by **date**, **type**, and **category**;
- Generate **monthly and yearly reports**.

---

## 🧱 Features

| Category | Description |
|-----------|--------------|
| 💰 **Transactions** | Create, list, update, and delete financial transactions |
| 📊 **Reports** | Retrieve total balance, category summary, and period summary |
| 🗓️ **Filters** | Filter by type (`income` or `expense`), category, and date range |
| 🧾 **Categories** | Full CRUD for custom categories |
| 🔒 **Authentication (extra)** | *(Optional challenge)* Implement JWT authentication and user control |

---

## ⚙️ Tech Stack

- **Go 1.22+**
- **MariaDB 10+**
- **Gin Gonic** — Web framework
- **GORM** — ORM for Go
- **Docker / Docker Compose**
- **godotenv** — Environment configuration
- **Swagger** — API documentation

---

## 📁 Project Structure

```
gofinance/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── connection.go
│   ├── models/
│   │   ├── transaction.go
│   │   └── category.go
│   ├── repository/
│   │   ├── transaction_repository.go
│   │   └── category_repository.go
│   ├── service/
│   │   ├── transaction_service.go
│   │   └── category_service.go
│   └── handlers/
│       ├── transaction_handler.go
│       └── category_handler.go
├── routes/
│   └── routes.go
├── Dockerfile
├── docker-compose.yml
├── .env
├── go.mod
├── go.sum
└── README.md
```

---

## ⚡ Environment Setup

### 1️⃣ Clone the repository
```bash
git clone https://github.com/yourusername/gofinance.git
cd gofinance
```

### 2️⃣ Create the `.env` file
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=secret
DB_NAME=gofinance
APP_PORT=8080
```

### 3️⃣ Start the environment with Docker
```bash
docker-compose up -d
```

### 4️⃣ Run the application
```bash
go run cmd/main.go
```

The API will be available at:  
👉 `http://localhost:8080/api`

---

## 🧩 Example Endpoints

### 🔹 Transactions
| Method | Endpoint | Description |
|--------|-----------|-------------|
| `POST` | `/api/transactions` | Create a new transaction |
| `GET` | `/api/transactions` | List all transactions |
| `GET` | `/api/transactions/:id` | Get transaction by ID |
| `PUT` | `/api/transactions/:id` | Update a transaction |
| `DELETE` | `/api/transactions/:id` | Delete a transaction |

### 🔹 Categories
| Method | Endpoint | Description |
|--------|-----------|-------------|
| `POST` | `/api/categories` | Create a new category |
| `GET` | `/api/categories` | List all categories |
| `PUT` | `/api/categories/:id` | Update a category |
| `DELETE` | `/api/categories/:id` | Delete a category |

### 🔹 Reports
| Method | Endpoint | Description |
|--------|-----------|-------------|
| `GET` | `/api/reports/summary` | Get total balance and general summary |
| `GET` | `/api/reports/monthly?month=10&year=2025` | Get detailed monthly report |

---

## 🧠 Extra Challenges

- Implement **JWT authentication** and multiple users.  
- Add **unit and integration tests**.  
- Build **advanced filters** (date ranges, multiple categories).  
- Add **asynchronous routines** (e.g., sending monthly reports via goroutines).  
- Implement **Redis cache** for reports.

---

## 🧑‍💻 Author

**Nathan Tanzi**  
[GitHub](https://github.com/yourusername) • [LinkedIn](https://linkedin.com/in/yourlinkedin)

---

## 📜 License

This project is licensed under the [MIT License](LICENSE).
