package repository

import (
	"database/sql"

	"github.com/Ntanzi07/gofinance/internal/models"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// GetAllTransactions calls the stored procedure `GetAllTransactions` and
// returns a list of transactions joined with user information.
func (r *TransactionRepository) GetAllTransactions() ([]models.TransactionWithUser, error) {
	rows, err := r.DB.Query("Call GetAllTransactions()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.TransactionWithUser
	for rows.Next() {
		var t models.TransactionWithUser
		if err := rows.Scan(&t.ID, &t.Type, &t.Amount, &t.Description, &t.Date, &t.UserName, &t.UserEmail); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)

	}
	return transactions, nil
}

// GetTransactionByID calls `GetTransactionById` stored procedure and returns
// a single transaction (with user info) for the provided ID.
func (r *TransactionRepository) GetTransactionByID(id int) (models.TransactionWithUser, error) {
	var t models.TransactionWithUser
	err := r.DB.QueryRow("CALL GetTransactionById(?)", id).Scan(&t.ID, &t.Type, &t.Amount, &t.Description, &t.Date, &t.UserName, &t.UserEmail)
	if err != nil {
		return models.TransactionWithUser{}, err
	}
	return t, nil
}
