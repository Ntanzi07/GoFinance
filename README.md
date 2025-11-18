# GoFinance — Financial Control API (current state)

This repository contains a REST API written in Go to manage personal financial transactions, backed by MariaDB. The documentation below describes how the current project is organized and how it works in this repository: the project uses `fiber` as HTTP framework, `database/sql` with stored procedures for data access, and a `handlers` / `repository` organization.

---

## Objective

The API provides features to:
- Create, read, update and delete transactions (income and expenses);
- Filter transactions by user and other criteria;
- Protect routes with JWT authentication where needed.

---

## Implemented features (current)

- Transaction CRUD implemented via repository methods that call stored procedures.
- User-scoped routes using the `/:name` route prefix.
- JWT-based authentication using the Fiber JWT middleware.
- `*sql.DB` is injected into repositories (no global DB variables).

---

## Current tech stack

- Go 1.22+
- MariaDB / MySQL (compatible)
- `github.com/gofiber/fiber/v2` — web framework
- `database/sql` — standard database access (no ORM)
- MySQL driver (e.g. `github.com/go-sql-driver/mysql`)
- `github.com/gofiber/contrib/jwt` — JWT middleware
- Docker / Docker Compose for local DB (optional)

Note: some queries rely on stored procedures such as `GetAllTransactions` and `GetTransactionById`. Ensure those procedures exist in your database before running the app.

---

## Project layout

```
GoFinance/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── db_config.go
│   ├── database/
│   │   └── db_connection.go
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── transaction_handler.go
│   │   └── user_handler.go
│   ├── models/
│   │   ├── db_model.go
│   │   ├── transactions_model.go
│   │   └── user_model.go
│   ├── repository/
│   │   ├── transactions_repository.go
│   │   └── user_repository.go
│   └── routes/
│       ├── routes.go
│       ├── routes_transaction.go
│       └── routes_user.go
├── docs/
│   └── swagger.yaml
├── docker-compose.yml
├── start-docker.sh
├── .env
├── go.mod
└── README.md
```

Key files present in the repository:
- `internal/repository/transactions_repository.go` — repository using `*sql.DB` and stored procedures
- `internal/handlers/transaction_handler.go` — handlers that receive repositories by dependency injection
- `internal/routes/routes_user.go` — user routes protected by JWT middleware

---

## Running locally

1. Create a `.env` file with your database and app settings, for example:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=secret
DB_NAME=gofinance
APP_PORT=8080
JWT_SECRET=your_jwt_secret
```

2. Start the database (if using Docker Compose):
```bash
docker-compose up -d
```

3. Run the application:
```bash
go run cmd/main.go
```

The API will be available at `http://localhost:8080` (or the value set in `APP_PORT`).

---

## Main endpoints (examples)

Transactions (admin / permission-protected):
- `GET  /transactions` → list all transactions (admin check applied)
- `GET  /transactions/:id` → get transaction by ID

User-scoped routes (group protected by JWT with prefix `/:name`):
- `GET  /:name/infos` → get user information
- `GET  /:name/` → list transactions belonging to `:name`
- `POST /:name/` → create a transaction for the user
- `PUT  /:name/transactions/:id` → update a user's transaction
- `DELETE /:name/transactions/:id` → delete a user's transaction

Example request to list user transactions:
```
GET /joao/transactions
Authorization: Bearer <token>
```

---

## Authentication

- JWT tokens are validated by middleware (`github.com/gofiber/contrib/jwt`).
- Handlers may check token claims (for example, an `isAdmin` claim) to authorize administrative actions.

---

## Notes and best practices

- The `*sql.DB` connection is injected into repositories with constructors like `NewTransactionRepository(db)`, and handlers receive repositories via `NewTransactionHandler(repo)`. This approach avoids global variables and improves testability.
- For unit tests, create mock implementations of the repository interfaces and inject them into handlers.
- Ensure the stored procedures the code expects are present in the database before running the application.

---

## Next suggested steps

- Add Swagger/OpenAPI documentation and optionally expose `/docs`.
- Add unit and integration tests for handlers and repositories.
- Add database migrations and SQL scripts (e.g. a `migrations/` folder or use `golang-migrate`).

---

##  Author

**Nathan Tanzi**  
[GitHub](https://github.com/Ntanzi07) • [LinkedIn](https://linkedin.com/in/nathan-tanzi)

---

## 📜 License

This project is licensed under the [MIT License](LICENSE).
