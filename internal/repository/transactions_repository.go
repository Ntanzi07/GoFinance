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

func (r *TransactionRepository) GetTransactionByID(id int) (models.TransactionWithUser, error) {
	var t models.TransactionWithUser
	err := r.DB.QueryRow("CALL GetTransactionById(?)", id).Scan(&t.ID, &t.Type, &t.Amount, &t.Description, &t.Date, &t.UserName, &t.UserEmail)
	if err != nil {
		return models.TransactionWithUser{}, err
	}
	return t, nil
}

func (r *TransactionRepository) CreateTransaction(userID int, tType string, amount float64, description, date string) error {
	_, err := r.DB.Exec("CALL CreateTransaction(?,?,?,?,?)", userID, tType, amount, description, date)
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) DeleteTransaction(transactionID int) error {
	_, err := r.DB.Exec("CALL DeleteTransaction(?)", transactionID)
	if err != nil {
		return err
	}
	return nil
}
